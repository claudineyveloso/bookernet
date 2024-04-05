package bucket

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	bucketStore types.BucketStore
}

func NewHandler(bucketStore types.BucketStore) *Handler {
	return &Handler{bucketStore: bucketStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_bucket", h.handleCreateBucket).Methods(http.MethodPost)
	router.HandleFunc("/get_buckets", h.handleGetBuckets).Methods(http.MethodGet)
	router.HandleFunc("/get_bucket/{bucketID}", h.handleGetBucket).Methods(http.MethodGet)

}

func (h *Handler) handleCreateBucket(w http.ResponseWriter, r *http.Request) {
	var bucket types.CreateBucketPayload
	if err := utils.ParseJSON(r, &bucket); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(bucket); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Payload inválido: %v", errors))
		return
	}
	err := h.bucketStore.CreateBucket(bucket)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, bucket)

}

func (h *Handler) handleGetBuckets(w http.ResponseWriter, r *http.Request) {
	buckets, err := h.bucketStore.GetBuckets()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter o Bucket: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, buckets)
}

func (h *Handler) handleGetBucket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["bucketID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Bucket ausente!"))
		return
	}
	parsedBucketsID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Bucket inválido!"))
		return
	}

	bucket, err := h.bucketStore.GetBucketByID(parsedBucketsID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, bucket)
}
