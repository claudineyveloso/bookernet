package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/google/uuid"
)

type OwnerRequest struct {
	ID         uuid.UUID      `json:"id"`
	PeopleType string         `json:"people_type"`
	IsActive   bool           `json:"is_active"`
	BucketID   uuid.UUID      `json:"bucket_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	Person     PersonRequest  `json:"person"`
	Address    AddressRequest `json:"address"`
}

func CreateOwnerController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var createOwnerRequest OwnerRequest
	if err := json.NewDecoder(r.Body).Decode(&createOwnerRequest); err != nil {
		fmt.Println("Erro ao decodificar corpo da solicitação:", err)
		http.Error(w, "Erro ao decodificar corpo da solicitação", http.StatusBadRequest)
		return
	}

	//fmt.Println("createOwnerRequest.Person:", createOwnerRequest.Person)
	//fmt.Println("createOwnerRequest.Address:", createOwnerRequest.Address)

	createOwnerRequest.ID = uuid.New()
	now := time.Now()

	createOwnerParams := db.CreateOwnerParams{
		ID:         createOwnerRequest.ID,
		PeopleType: createOwnerRequest.PeopleType,
		IsActive:   createOwnerRequest.IsActive,
		BucketID:   createOwnerRequest.BucketID,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	createPersonParams := db.CreatePersonParams{
		ID:             uuid.New(),
		FirstName:      createOwnerRequest.Person.FirstName,
		LastName:       createOwnerRequest.Person.LastName,
		Email:          createOwnerRequest.Person.Email,
		Phone:          utils.CreateNullString(createOwnerRequest.Person.Phone),
		CellPhone:      createOwnerRequest.Person.CellPhone,
		PersonableID:   createOwnerRequest.ID,
		PersonableType: "owner",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	createAddressParams := db.CreateAddressParams{
		ID:              uuid.New(),
		PublicPlace:     utils.CreateNullString(createOwnerRequest.Address.PublicPlace),
		Complement:      utils.CreateNullString(createOwnerRequest.Address.Complement),
		Neighborhood:    utils.CreateNullString(createOwnerRequest.Address.Neighborhood),
		City:            utils.CreateNullString(createOwnerRequest.Address.City),
		State:           utils.CreateNullString(createOwnerRequest.Address.State),
		ZipCode:         utils.CreateNullString(createOwnerRequest.Address.ZipCode),
		AddressableID:   createOwnerRequest.ID,
		AddressableType: "owner",
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	if err := queries.CreateOwner(r.Context(), createOwnerParams); err != nil {
		http.Error(w, "Erro ao criar proprietário", http.StatusBadRequest)
		return
	}

	// Create person of owner
	if err := queries.CreatePerson(r.Context(), createPersonParams); err != nil {
		http.Error(w, "Erro ao criar a pessoa do proprietário", http.StatusBadRequest)
		return
	}

	// Create address of owner
	if err := queries.CreateAddress(r.Context(), createAddressParams); err != nil {
		http.Error(w, "Erro ao criar o endereço do proprietário", http.StatusBadRequest)
		return
	}

	// Construir JSON de resposta
	response := struct {
		OwnerID    uuid.UUID      `json:"owner_id"`
		PeopleType string         `json:"people_type"`
		IsActive   bool           `json:"is_active"`
		BucketID   uuid.UUID      `json:"bucket_id"`
		Person     PersonRequest  `json:"person"`
		Address    AddressRequest `json:"address"`
		CreatedAt  time.Time      `json:"created_at"`
		UpdatedAt  time.Time      `json:"updated_at"`
	}{
		OwnerID:    createOwnerRequest.ID,
		PeopleType: createOwnerRequest.PeopleType,
		IsActive:   createOwnerRequest.IsActive,
		BucketID:   createOwnerRequest.BucketID,
		Person:     createOwnerRequest.Person,
		Address:    createOwnerRequest.Address,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
