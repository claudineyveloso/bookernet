package insurance

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
	insuranceStore types.InsuranceStore
}

func NewHandler(insuranceStore types.InsuranceStore) *Handler {
	return &Handler{insuranceStore: insuranceStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_insurance", h.handleCreateInsurance).Methods(http.MethodPost)
	router.HandleFunc("/get_insurances", h.handleGetInsurances).Methods(http.MethodGet)
	//router.HandleFunc("/get_insurances", auth.WithJWTAuth(h.handleGetBuckets, h.insuranceStore)).Methods(http.MethodPost)
	router.HandleFunc("/get_insurance/{insuranceID}", h.handleGetInsurance).Methods(http.MethodGet)

}

func (h *Handler) handleCreateInsurance(w http.ResponseWriter, r *http.Request) {
	var insurance types.InsurancePayload
	if err := utils.ParseJSON(r, &insurance); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(insurance); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Payload inválido: %v", errors))
		return
	}
	err := h.insuranceStore.CreateInsurance(insurance)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, insurance)

}

func (h *Handler) handleGetInsurances(w http.ResponseWriter, r *http.Request) {
	//insuranceID := auth.GetUserIDFromContext(r.Context())
	//fmt.Println("Valor de userID", insuranceID)
	insurances, err := h.insuranceStore.GetInsurances()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter o Convênio: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, insurances)
}

func (h *Handler) handleGetInsurance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["insuranceID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Convênio ausente!"))
		return
	}
	parsedInsurancesID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Convênio inválido!"))
		return
	}

	insurance, err := h.insuranceStore.GetInsuranceByID(parsedInsurancesID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, insurance)
}
