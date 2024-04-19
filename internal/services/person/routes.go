package person

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	personStore types.PersonStore
}

func NewHandler(personStore types.PersonStore) *Handler {
	return &Handler{personStore: personStore}
}

func (h *Handler) handleCreatePerson(w http.ResponseWriter, r *http.Request) {
	var person types.PersonPayload
	if err := utils.ParseJSON(r, &person); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(person); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Payload inválido: %v", errors))
		return
	}
	err := h.personStore.CreatePerson(person)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, person)

}
