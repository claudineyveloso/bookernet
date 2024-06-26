package owner

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
	ownerStore   types.OwnerStore
	personStore  types.PersonStore
	addressStore types.AddressStore
}

func NewHandler(ownerStore types.OwnerStore, personStore types.PersonStore, addressStore types.AddressStore) *Handler {
	return &Handler{ownerStore: ownerStore, personStore: personStore, addressStore: addressStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create_owner", h.handleCreateOwner).Methods(http.MethodPost)
	router.HandleFunc("/get_owners", h.handleGetOwners).Methods(http.MethodGet)
	router.HandleFunc("/get_owner/{id}", h.handleGetOwner).Methods(http.MethodGet)
}

func (h *Handler) handleCreateOwner(w http.ResponseWriter, r *http.Request) {
	var owner types.OwnerPayload
	if err := utils.ParseJSON(r, &owner); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(owner); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Payload inválido: %v", errors))
		return
	}
	createdOwner, err := h.ownerStore.CreateOwner(owner)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	person := types.PersonPayload{
		FirstName:      owner.Person.FirstName,
		LastName:       owner.Person.LastName,
		Email:          owner.Person.Email,
		Phone:          owner.Person.Phone,
		CellPhone:      owner.Person.CellPhone,
		PersonableID:   createdOwner,
		PersonableType: "owner",
	}
	if err := h.personStore.CreatePerson(person); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	address := types.AddressPayload{
		PublicPlace:     owner.Address.PublicPlace,
		Complement:      owner.Address.Complement,
		Neighborhood:    owner.Address.Neighborhood,
		City:            owner.Address.City,
		State:           owner.Address.State,
		ZipCode:         owner.Address.ZipCode,
		AddressableID:   createdOwner,
		AddressableType: "owner",
	}

	if err := h.addressStore.CreateAddress(address); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, owner)

}

func (h *Handler) handleGetOwners(w http.ResponseWriter, r *http.Request) {
	owners, err := h.ownerStore.GetOwners()
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter o Proprietário: %v", err), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, owners)
}

func (h *Handler) handleGetOwner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do proprietário ausente!"))
		return
	}
	parsedID, err := uuid.Parse(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ID do proprietário inválido!"))
		return
	}

	user, err := h.ownerStore.GetOwner(parsedID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, user)
}
