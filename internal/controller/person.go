package controller

import (
	"net/http"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PersonRequest struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name" validate:"required"`
	LastName       string    `json:"last_name" validate:"required"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	CellPhone      string    `json:"cell_phone" validate:"required"`
	PersonableID   uuid.UUID `json:"personable_id"`
	PersonableType string    `json:"personable_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (o *PersonRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(o); err != nil {
		return err
	}
	return nil
}
func CreatePersonController(w http.ResponseWriter, r *http.Request, queries *db.Queries, createPersonRequest *OwnerRequest) (uuid.UUID, error) {
	// Validate user fields before continuing
	//
	if err := createPersonRequest.Person.Validate(); err != nil {
		// Se houver erros de validação, tratamos o erro aqui
		errorMsg := "Erros de validação:"
		for _, err := range err.(validator.ValidationErrors) {
			errorMsg += " " + err.Field() + " está em branco;"
		}
		http.Error(w, errorMsg, http.StatusBadRequest)
		return uuid.Nil, err
	}

	now := time.Now()

	createPersonParams := db.CreatePersonParams{
		ID:             uuid.New(),
		FirstName:      createPersonRequest.Person.FirstName,
		LastName:       createPersonRequest.Person.LastName,
		Email:          createPersonRequest.Person.Email,
		Phone:          utils.CreateNullString(createPersonRequest.Person.Phone),
		CellPhone:      createPersonRequest.Person.CellPhone,
		PersonableID:   createPersonRequest.ID,
		PersonableType: "owner",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	err := queries.CreatePerson(r.Context(), createPersonParams)
	if err != nil {
		return uuid.Nil, err
	}

	return createPersonParams.ID, nil
}
