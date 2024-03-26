package owner

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	ownerStore types.OwnerStore
}

func NewHandler(ownerStore types.OwnerStore) *Handler {
	return &Handler{ownerStore: ownerStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_owner", h.handleCreateOwner).Methods(http.MethodPost)
}

func (h *Handler) handleCreateOwner(w http.ResponseWriter, r *http.Request) {
	var owner types.CreateOwnerPayload
	if err := utils.ParseJSON(r, &owner); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(owner); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	err := h.ownerStore.CreateOwner(owner)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, owner)

}
