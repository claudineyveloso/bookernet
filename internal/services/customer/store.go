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
		fmt.Println("Erro ao criar o endereço:", err)
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
	for _, dbUser := range dbCustomers {
		customer := convertDBCustomerToTypeCustomer(dbUser)
		customers = append(customers, customer)
	}
	return customers, nil
}

// func (s *Store) GetCustomersWithDetails() ([]*types.GetCustomersRow, error) {
// 	queries := db.New(s.db)
// 	ctx := context.Background()

// 	dbCustomers, err := queries.GetCustomersWithDetails(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var customers []*types.GetCustomersRow
// 	for _, dbUser := range dbCustomers {
// 		customer := convertDBCustomerToTypeCustomer(dbUser)
// 		getCustomerRow := convertCustomerToGetCustomersRow(*customer) // Desreferenciando o ponteiro aqui
// 		customers = append(customers, getCustomerRow)
// 	}
// 	return customers, nil
// }

// func convertCustomerToGetCustomersRow(dbCustomer types.Customer) *types.GetCustomersRow {
// 	customer := &types.GetCustomersRow{
// 		ID: dbCustomer.ID,
// 		//Birthday:        dbUser.Birthday,
// 		CreatedAt: dbCustomer.CreatedAt,
// 		UpdatedAt: dbCustomer.UpdatedAt,
// 		ID_2:      dbCustomer.Person.ID,
// 		FirstName: dbCustomer.Person.FirstName,
// 		LastName:  dbCustomer.Person.LastName,
// 		Email:     dbCustomer.Person.Email,
// 		//Phone:           dbUser.Person.Phone,
// 		CellPhone:       dbCustomer.Person.CellPhone,
// 		PersonableID:    dbCustomer.Person.PersonableID,
// 		PersonableType:  dbCustomer.Person.PersonableType,
// 		CreatedAt_2:     dbCustomer.Person.CreatedAt,
// 		UpdatedAt_2:     dbCustomer.Person.UpdatedAt,
// 		ID_3:            dbCustomer.Address.ID,
// 		PublicPlace:     dbCustomer.Address.PublicPlace,
// 		Complement:      dbCustomer.Address.Complement,
// 		Neighborhood:    dbCustomer.Address.Neighborhood,
// 		City:            dbCustomer.Address.City,
// 		State:           dbCustomer.Address.State,
// 		ZipCode:         dbCustomer.Address.ZipCode,
// 		AddressableID:   dbCustomer.Address.AddressableID,
// 		AddressableType: dbCustomer.Address.AddressableType,
// 		CreatedAt_3:     dbCustomer.Address.CreatedAt,
// 		UpdatedAt_3:     dbCustomer.Address.UpdatedAt,
// 	}
// 	return customer
// }

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
