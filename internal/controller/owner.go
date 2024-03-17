package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type OwnerRequest struct {
	ID         uuid.UUID      `json:"id"`
	PeopleType string         `json:"people_type" validate:"required"`
	IsActive   bool           `json:"is_active" validate:"required"`
	BucketID   uuid.UUID      `json:"bucket_id" validate:"required"`
	Person     PersonRequest  `json:"person"`
	Address    AddressRequest `json:"address"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func CreateOwnerController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var createOwnerRequest OwnerRequest
	if err := json.NewDecoder(r.Body).Decode(&createOwnerRequest); err != nil {
		http.Error(w, "Erro ao decodificar corpo da solicitação", http.StatusBadRequest)
		return
	}

	// Validate user fields before continuing
	if err := createOwnerRequest.Validate(); err != nil {
		http.Error(w, "Erro de validação: "+err.Error(), http.StatusBadRequest)
		return
	}

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

	if err := queries.CreateOwner(r.Context(), createOwnerParams); err != nil {
		http.Error(w, "Erro ao criar proprietário", http.StatusBadRequest)
		return
	}
	// CreatePersonController creates a new person in the database based on the data provided in
	// createPersonRequest.
	_, err := CreatePersonController(w, r, queries, &createOwnerRequest)
	if err != nil {
		http.Error(w, "Erro ao criar a pessoa do proprietário", http.StatusBadRequest)
		return
	}

	// CreateAddressController creates a new address in the database based on the data
	// fornecidos em createAddressRequest.
	_, err = CreateAddressController(w, r, queries, &createOwnerRequest)
	if err != nil {
		http.Error(w, "Erro ao criar a pessoa do proprietário", http.StatusBadRequest)
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

func (o *OwnerRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(o)
	if err != nil {
		errorMap := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "PeopleType":
				errorMap["people_type"] = "O campo de tipo de usuário não pode estar vazio"
			case "IsActive":
				errorMap["is_active"] = "O campo proprietário ativo/inativo não pode estar vazio"
			case "BucketID":
				errorMap["bucket_id"] = "O bucket do proprietário não pode estar vazio"
			}
		}
		var errorMsg string
		for _, msg := range errorMap {
			errorMsg += msg + "\n"
		}

		return errors.New(errorMsg)
	}
	return nil
}
