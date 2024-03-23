package user

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	userStore types.UserStore
}

func NewHandler(userStore types.UserStore) *Handler {
	return &Handler{userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/get_users", h.handleGetUsers).Methods(http.MethodGet)
}

func (h *Handler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userStore.GetUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter usuários: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, users)
}
