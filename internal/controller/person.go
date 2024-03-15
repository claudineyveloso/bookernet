package controller

import (
	"net/http"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/google/uuid"
)

type PersonRequest struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	CellPhone      string    `json:"cell_phone"`
	PersonableID   uuid.UUID `json:"personable_id"`
	PersonableType string    `json:"personable_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func CreatePersonController(w http.ResponseWriter, r *http.Request, queries *db.Queries, createPersonRequest *OwnerRequest) (uuid.UUID, error) {
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
