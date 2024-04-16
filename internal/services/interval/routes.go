package interval

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/services/auth"
	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	intervalStore types.IntervalStore
}

func NewHandler(intervalStore types.IntervalStore) *Handler {
	return &Handler{intervalStore: intervalStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_interval", h.handleCreateInterval).Methods(http.MethodPost)
	router.HandleFunc("/get_intervals", h.handleGetIntervals).Methods(http.MethodGet)
	//router.HandleFunc("/get_buckets", auth.WithJWTAuth(h.handleGetBuckets, h.userStore)).Methods(http.MethodPost)
	router.HandleFunc("/get_interval/{intervalID}", h.handleGetInterval).Methods(http.MethodGet)

}

func (h *Handler) handleCreateInterval(w http.ResponseWriter, r *http.Request) {
	var interval types.CreateIntervalPayload
	if err := utils.ParseJSON(r, &interval); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(interval); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Payload inválido: %v", errors))
		return
	}
	err := h.intervalStore.CreateInterval(interval)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, interval)

}

func (h *Handler) handleGetIntervals(w http.ResponseWriter, r *http.Request) {
	bucketID := auth.GetUserIDFromContext(r.Context())
	fmt.Println("Valor de userID", bucketID)
	intervals, err := h.intervalStore.GetIntervals()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter o Intervalo: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, intervals)
}

func (h *Handler) handleGetInterval(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["intervalID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Intervalo ausente!"))
		return
	}
	parsedIntervalsID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Intervalo inválido!"))
		return
	}

	interval, err := h.intervalStore.GetIntervalByID(parsedIntervalsID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, interval)
}
