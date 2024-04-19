package customer

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
	customerStore types.CustomerStore
	personStore   types.PersonStore
	addressStore  types.AddressStore
}

func NewHandler(customerStore types.CustomerStore, personStore types.PersonStore, addressStore types.AddressStore) *Handler {
	return &Handler{customerStore: customerStore, personStore: personStore, addressStore: addressStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_customer", h.handleCreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/get_customers", h.handleGetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/get_customer/{id}", h.handleGetCustomer).Methods(http.MethodGet)
}

func (h *Handler) handleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer types.CustomerPayload
	if err := utils.ParseJSON(r, &customer); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(customer); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Payload inválido: %v", errors))
		return
	}
	createdCustomer, err := h.customerStore.CreateCustomer(customer)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	person := types.PersonPayload{
		FirstName:      customer.Person.FirstName,
		LastName:       customer.Person.LastName,
		Email:          customer.Person.Email,
		Phone:          customer.Person.Phone,
		CellPhone:      customer.Person.CellPhone,
		PersonableID:   createdCustomer,
		PersonableType: "customer",
	}

	if err := h.personStore.CreatePerson(person); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	address := types.AddressPayload{
		PublicPlace:     customer.Address.PublicPlace,
		Complement:      customer.Address.Complement,
		Neighborhood:    customer.Address.Neighborhood,
		City:            customer.Address.City,
		State:           customer.Address.State,
		ZipCode:         customer.Address.ZipCode,
		AddressableID:   createdCustomer,
		AddressableType: "customer",
	}

	if err := h.addressStore.CreateAddress(address); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, customer)

}

func (h *Handler) handleGetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerStore.GetCustomers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter o Cliente: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, customers)
}

func (h *Handler) handleGetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Cliente ausente!"))
		return
	}
	parsedID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do Cliente inválido!"))
		return
	}

	user, err := h.customerStore.GetCustomer(parsedID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}
