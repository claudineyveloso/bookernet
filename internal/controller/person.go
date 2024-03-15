package controller

import (
	"encoding/json"
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

func CreatePersonController(w http.ResponseWriter, r *http.Request, queries *db.Queries, ownerID uuid.UUID) error {
	var createPersonRequest PersonRequest
	if err := json.NewDecoder(r.Body).Decode(&createPersonRequest); err != nil {
		http.Error(w, "Erro ao decodificar corpo da solicitação"+err.Error(), http.StatusBadRequest)
		return err
	}

	now := time.Now()
	createPersonParams := CreatePersonParamsFromRequest(createPersonRequest, ownerID, now)

	if err := queries.CreatePerson(r.Context(), createPersonParams); err != nil {
		http.Error(w, "Erro ao criar a pessoa: %v"+err.Error(), http.StatusBadRequest)
		return err

	}

	return nil
}

func CreatePersonParamsFromRequest(req PersonRequest, ownerID uuid.UUID, now time.Time) db.CreatePersonParams {
	return db.CreatePersonParams{
		ID:             uuid.New(),
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          utils.CreateNullString(req.Phone),
		CellPhone:      req.CellPhone,
		PersonableID:   ownerID,
		PersonableType: "owner",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
