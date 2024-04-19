package owner

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

func (s *Store) CreateOwner(owner types.OwnerPayload) (uuid.UUID, error) {
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
	if err := queries.CreateOwner(ctx, createOwnerParams); err != nil {
		fmt.Println("Erro ao criar o Proprietário:", err)
		return uuid.UUID{}, err
	}

	return owner.ID, nil
}

func (s *Store) GetOwners() ([]*types.Owner, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbOwners, err := queries.GetOwnersWithDetails(ctx)
	if err != nil {
		return nil, err
	}

	var owners []*types.Owner
	for _, dbOwner := range dbOwners {
		owner := convertDBOwnerToTypeOwner(dbOwner)
		owners = append(owners, owner)
	}
	return owners, nil
}

func (s *Store) GetOwner(id uuid.UUID) (*types.Owner, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbOwner, err := queries.GetOwnerWithDetails(ctx, id)
	if err != nil {
		return nil, err
	}

	owner := convertDBOwnerToTypeOwner(dbOwner)
	return owner, nil
}

func convertDBOwnerToTypeOwner(dbOwner db.FullOwnerRow) *types.Owner {
	owner := &types.Owner{
		ID:         dbOwner.Owner.ID,
		PeopleType: dbOwner.Owner.PeopleType,
		IsActive:   dbOwner.Owner.IsActive,
		BucketID:   dbOwner.Owner.BucketID,
		CreatedAt:  dbOwner.Owner.CreatedAt,
		UpdatedAt:  dbOwner.Owner.UpdatedAt,
		Person: types.Person{
			ID:             dbOwner.Person.ID,
			FirstName:      dbOwner.Person.FirstName,
			LastName:       dbOwner.Person.LastName,
			Email:          dbOwner.Person.Email,
			Phone:          utils.GetValidString(dbOwner.Person.Phone),
			CellPhone:      dbOwner.Person.CellPhone,
			PersonableID:   dbOwner.Person.PersonableID,
			PersonableType: dbOwner.Person.PersonableType,
			CreatedAt:      dbOwner.Person.CreatedAt,
			UpdatedAt:      dbOwner.Person.UpdatedAt,
		},
		Address: types.Address{
			ID:              dbOwner.Address.ID,
			PublicPlace:     utils.GetValidString(dbOwner.Address.PublicPlace),
			Complement:      utils.GetValidString(dbOwner.Address.Complement),
			Neighborhood:    utils.GetValidString(dbOwner.Address.Neighborhood),
			City:            utils.GetValidString(dbOwner.Address.City),
			State:           utils.GetValidString(dbOwner.Address.State),
			ZipCode:         utils.GetValidString(dbOwner.Address.ZipCode),
			AddressableID:   dbOwner.Address.AddressableID,
			AddressableType: dbOwner.Address.AddressableType,
			CreatedAt:       dbOwner.Address.CreatedAt,
			UpdatedAt:       dbOwner.Address.UpdatedAt,
		},
	}
	return owner
}
