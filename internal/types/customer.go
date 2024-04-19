package types

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `json:"id"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Person    Person    `json:"person"`
	Address   Address   `json:"address"`
}

type GetCustomersRow struct {
	ID              uuid.UUID    `json:"id"`
	Birthday        sql.NullTime `json:"birthday"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	ID_2            uuid.UUID    `json:"id_2"`
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name"`
	Email           string       `json:"email"`
	Phone           string       `json:"phone"`
	CellPhone       string       `json:"cell_phone"`
	PersonableID    uuid.UUID    `json:"personable_id"`
	PersonableType  string       `json:"personable_type"`
	CreatedAt_2     time.Time    `json:"created_at_2"`
	UpdatedAt_2     time.Time    `json:"updated_at_2"`
	ID_3            uuid.UUID    `json:"id_3"`
	PublicPlace     string       `json:"public_place"`
	Complement      string       `json:"complement"`
	Neighborhood    string       `json:"neighborhood"`
	City            string       `json:"city"`
	State           string       `json:"state"`
	ZipCode         string       `json:"zip_code"`
	AddressableID   uuid.UUID    `json:"addressable_id"`
	AddressableType string       `json:"addressable_type"`
	CreatedAt_3     time.Time    `json:"created_at_3"`
	UpdatedAt_3     time.Time    `json:"updated_at_3"`
}

type CustomerPayload struct {
	ID        uuid.UUID `json:"id"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Person    Person    `json:"person"`
	Address   Address   `json:"address"`
}

type CustomerStore interface {
	CreateCustomer(CustomerPayload) (uuid.UUID, error)
	GetCustomers() ([]*Customer, error)
	GetCustomer(id uuid.UUID) (*Customer, error)
}
