package person

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/claudineyveloso/bookernet.git/internal/utils"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreatePerson(person types.CreatePersonPayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	person.ID = uuid.New()
	now := time.Now()
	person.CreatedAt = now
	person.UpdatedAt = now

	createPersonParams := db.CreatePersonParams{
		ID:             person.ID,
		FirstName:      person.FirstName,
		LastName:       person.LastName,
		Email:          person.Email,
		Phone:          utils.CreateNullString(person.Phone),
		CellPhone:      person.CellPhone,
		PersonableID:   person.PersonableID,
		PersonableType: person.PersonableType,
		CreatedAt:      person.CreatedAt,
		UpdatedAt:      person.UpdatedAt,
	}
	if err := queries.CreatePerson(ctx, createPersonParams); err != nil {
		fmt.Println("Erro ao criar a pessoa:", err)
		return err
	}
	return nil
}
