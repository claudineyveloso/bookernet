package owner

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/types"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOwner(owner types.CreateOwnerPayload) (uuid.UUID, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	owner.ID = uuid.New()
	now := time.Now()
	owner.CreatedAt = now
	owner.UpdatedAt = now

	createOwnerParams := db.CreateOwnerParams{
		ID:         owner.ID,
		PeopleType: owner.PeopleType,
		IsActive:   owner.IsActive,
		BucketID:   owner.BucketID,
		CreatedAt:  owner.CreatedAt,
		UpdatedAt:  owner.UpdatedAt,
	}
	fmt.Println(createOwnerParams)
	if err := queries.CreateOwner(ctx, createOwnerParams); err != nil {
		fmt.Println("Erro ao criar o endereço:", err)
		return uuid.UUID{}, err
	}

	return owner.ID, nil
}
