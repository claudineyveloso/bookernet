package customer

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

func (s *Store) CreateCustomer(customer types.CreateCustomerPayload) (uuid.UUID, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	customer.ID = uuid.New()
	now := time.Now()
	customer.CreatedAt = now
	customer.UpdatedAt = now

	createCustomerParams := db.CreateCustomerParams{
		ID:        customer.ID,
		Birthday:  utils.CreateNullDate(customer.Birthday),
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}

	if err := queries.CreateCustomer(ctx, createCustomerParams); err != nil {
		fmt.Println("Erro ao criar o endereço:", err)
		return uuid.UUID{}, err
	}
	return customer.ID, nil

}
