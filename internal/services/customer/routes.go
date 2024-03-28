package customer

import (
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/go-playground/validator/v10"
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
}

func (h *Handler) handleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer types.CreateCustomerPayload
	if err := utils.ParseJSON(r, &customer); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(customer); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	createdCustomer, err := h.customerStore.CreateCustomer(customer)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	person := types.CreatePersonPayload{
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

	address := types.CreateAddressPayload{
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
