package address

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

func (s *Store) CreateAddress(address types.CreateAddressPayload) error {
	queries := db.New(s.db)
	ctx := context.Background()

	address.ID = uuid.New()
	now := time.Now()
	address.CreatedAt = now
	address.UpdatedAt = now

	createAddressParams := db.CreateAddressParams{
		ID:              address.ID,
		PublicPlace:     utils.CreateNullString(address.PublicPlace),
		Complement:      utils.CreateNullString(address.Complement),
		Neighborhood:    utils.CreateNullString(address.Neighborhood),
		City:            utils.CreateNullString(address.City),
		State:           utils.CreateNullString(address.State),
		ZipCode:         utils.CreateNullString(address.ZipCode),
		AddressableID:   address.AddressableID,
		AddressableType: address.AddressableType,
		CreatedAt:       address.CreatedAt,
		UpdatedAt:       address.UpdatedAt,
	}
	if err := queries.CreateAddress(ctx, createAddressParams); err != nil {
		fmt.Println("Erro ao criar o endereço:", err)
		return err
	}
	return nil
}
