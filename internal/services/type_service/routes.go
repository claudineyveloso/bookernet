package typeservice

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
	typeServiceStore types.TypeServiceStore
}

func NewHandler(typeServiceStore types.TypeServiceStore) *Handler {
	return &Handler{typeServiceStore: typeServiceStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_type_service", h.handleCreateTypeService).Methods(http.MethodPost)
	router.HandleFunc("/get_type_services", h.handleGetTypeServices).Methods(http.MethodGet)
	router.HandleFunc("/get_type_service/{typeServiceID}", h.handleGetTypeService).Methods(http.MethodGet)

}

func (h *Handler) handleCreateTypeService(w http.ResponseWriter, r *http.Request) {
	var typeService types.TypeServicePayload
	if err := utils.ParseJSON(r, &typeService); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(typeService); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Payload inválido: %v", errors))
		return
	}
	err := h.typeServiceStore.CreateTypeService(typeService)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, typeService)

}

func (h *Handler) handleGetTypeServices(w http.ResponseWriter, r *http.Request) {
	typeServices, err := h.typeServiceStore.GetTypeServices()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter tipo de serviço: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, typeServices)
}

func (h *Handler) handleGetTypeService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["typeServiceID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Tipo de Serviço ausente!"))
		return
	}
	parsedTypeServicesID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Tipo de Serviço inválido!"))
		return
	}

	typeService, err := h.typeServiceStore.GetTypeServiceByID(parsedTypeServicesID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, typeService)
}
