package attendance

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
  attendanceStore types.AttendanceStore
}

func NewHandler(attendanceStore types.AttendanceStore) *Handler {
  return &Handler{attendanceStore: attendanceStore}
}


func (h *Handler) RegisterRoutes(routes *mux.Router) {
  routes.HandleFunc("/create_attendance", h.handleCreateAttendance).Methods(http.MethodPost)
  routes.HandleFunc("/get_attendances", h.handleGetAttendances).Methods(http.MethodGet)
  routes.HandleFunc("/get_attendance/{attendanceID}", h.handleGetAttendance).Methods(http.MethodGet)
  //routes.HandleFunc("/update_attendance/{attendanceID}", h.handleUpdateAttendance).Methods(http.MethodPut)
}

func (h *Handler) handleCreateAttendance(w http.ResponseWriter, r *http.Request) {
  var attendance types.CreateAttendancePayload
  if err := utils.ParseJSON(r, &attendance); err != nil {
    utils.WriteError(w, http.StatusBadRequest, err)
    return
  }
  if err := h.attendanceStore.CreateAttendance(attendance); err!= nil {
    utils.WriteError(w, http.StatusInternalServerError, err)
    return
  }
  utils.WriteJSON(w, http.StatusCreated, attendance)
}

func (h *Handler) handleGetAttendances(w http.ResponseWriter, r *http.Request) {
  attendances, err := h.attendanceStore.GetAttendances()
  if err != nil {
    http.Error(w, fmt.Sprintf("Erro ao obter uma confirmação de atendimento: %v", err), http.StatusInternalServerError)
    return
  }
  utils.WriteJSON(w, http.StatusOK, attendances)
}

func (h *Handler) handleGetAttendance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["attendanceID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID da confirmação de atendimento ausente!"))
		return
	}
	parsedAttendanceID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID da confirmação de atendimento inválido!"))
		return
	}

	attendance, err := h.attendanceStore.GetAttendanceByID(parsedAttendanceID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, attendance)
}

