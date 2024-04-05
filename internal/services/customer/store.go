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
		ID: customer.ID,
		//Birthday:  utils.CreateNullDate(customer.Birthday),
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}

	if err := queries.CreateCustomer(ctx, createCustomerParams); err != nil {
		fmt.Println("Erro ao criar o Cliente:", err)
		return uuid.UUID{}, err
	}
	return customer.ID, nil

}

func (s *Store) GetCustomers() ([]*types.Customer, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbCustomers, err := queries.GetCustomersWithDetails(ctx)
	if err != nil {
		return nil, err
	}

	var customers []*types.Customer
	for _, dbCustomer := range dbCustomers {
		customer := convertDBCustomerToTypeCustomer(dbCustomer)
		customers = append(customers, customer)
	}
	return customers, nil
}

func (s *Store) GetCustomer(id uuid.UUID) (*types.Customer, error) {
	queries := db.New(s.db)
	ctx := context.Background()

	dbCustomer, err := queries.GetCustomerWithDetails(ctx, id)
	if err != nil {
		return nil, err
	}

	customer := convertDBCustomerToTypeCustomer(dbCustomer)
	return customer, nil
}

func convertDBCustomerToTypeCustomer(dbCustomer db.FullCustomerRow) *types.Customer {
	customer := &types.Customer{
		ID: dbCustomer.Customer.ID,
		//Birthday:  dbCustomer.Customer.Birthday,
		CreatedAt: dbCustomer.Customer.CreatedAt,
		UpdatedAt: dbCustomer.Customer.UpdatedAt,
		Person: types.Person{
			ID:             dbCustomer.Person.ID,
			FirstName:      dbCustomer.Person.FirstName,
			LastName:       dbCustomer.Person.LastName,
			Email:          dbCustomer.Person.Email,
			Phone:          utils.GetValidString(dbCustomer.Person.Phone),
			CellPhone:      dbCustomer.Person.CellPhone,
			PersonableID:   dbCustomer.Person.PersonableID,
			PersonableType: dbCustomer.Person.PersonableType,
			CreatedAt:      dbCustomer.Person.CreatedAt,
			UpdatedAt:      dbCustomer.Person.UpdatedAt,
		},
		Address: types.Address{
			ID:              dbCustomer.Address.ID,
			PublicPlace:     utils.GetValidString(dbCustomer.Address.PublicPlace),
			Complement:      utils.GetValidString(dbCustomer.Address.Complement),
			Neighborhood:    utils.GetValidString(dbCustomer.Address.Neighborhood),
			City:            utils.GetValidString(dbCustomer.Address.City),
			State:           utils.GetValidString(dbCustomer.Address.State),
			ZipCode:         utils.GetValidString(dbCustomer.Address.ZipCode),
			AddressableID:   dbCustomer.Address.AddressableID,
			AddressableType: dbCustomer.Address.AddressableType,
			CreatedAt:       dbCustomer.Address.CreatedAt,
			UpdatedAt:       dbCustomer.Address.UpdatedAt,
		},
	}
	return customer
}
