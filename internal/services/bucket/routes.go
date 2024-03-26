package bucket

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	bucketStore types.BucketStore
}

func NewHandler(bucketStore types.BucketStore) *Handler {
	return &Handler{bucketStore: bucketStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	//router.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).Methods(http.MethodPost)
	router.HandleFunc("/create_bucket", h.handleCreateBucket).Methods(http.MethodPost)
	//router.HandleFunc("/get_users", h.handleGetUsers).Methods(http.MethodGet)
	//router.HandleFunc("/get_user/{userID}", h.handleGetUser).Methods(http.MethodGet)

}

func (h *Handler) handleCreateBucket(w http.ResponseWriter, r *http.Request) {
	var bucket types.CreateBucketPayload
	if err := utils.ParseJSON(r, &bucket); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(bucket); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	err := h.bucketStore.CreateBucket(bucket)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, bucket)

}

// func (h *Handler) handleGetUsers(w http.ResponseWriter, r *http.Request) {
// 	users, err := h.bucketStore.GetUsers()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Erro ao obter usuários: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	utils.WriteJSON(w, http.StatusOK, users)
// }

// func (h *Handler) handleGetUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	str, ok := vars["userID"]
// 	if !ok {
// 		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
// 		return
// 	}
// 	parsedUserID, err := uuid.Parse(str)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
// 		return
// 	}

// 	user, err := h.userStore.GetUserByID(parsedUserID)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	utils.WriteJSON(w, http.StatusOK, user)
// }
