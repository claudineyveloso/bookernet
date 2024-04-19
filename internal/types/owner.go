package types

import (
	"time"

	"github.com/google/uuid"
)

type Owner struct {
	ID         uuid.UUID `json:"id"`
	PeopleType string    `json:"people_type"`
	IsActive   bool      `json:"is_active"`
	BucketID   uuid.UUID `json:"bucket_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Person     Person    `json:"person"`
	Address    Address   `json:"address"`
}

type OwnerPayload struct {
	ID         uuid.UUID `json:"id"`
	PeopleType string    `json:"people_type" validate:"required"`
	IsActive   bool      `json:"is_active"`
	BucketID   uuid.UUID `json:"bucket_id" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Person     Person    `json:"person"`
	Address    Address   `json:"address"`
}

type GetOwnerRow struct {
	ID              uuid.UUID `json:"id"`
	PeopleType      string    `json:"people_type"`
	IsActive        bool      `json:"is_active"`
	BucketID        uuid.UUID `json:"bucket_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ID_2            uuid.UUID `json:"id_2"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	CellPhone       string    `json:"cell_phone"`
	PersonableID    uuid.UUID `json:"personable_id"`
	PersonableType  string    `json:"personable_type"`
	CreatedAt_2     time.Time `json:"created_at_2"`
	UpdatedAt_2     time.Time `json:"updated_at_2"`
	ID_3            uuid.UUID `json:"id_3"`
	PublicPlace     string    `json:"public_place"`
	Complement      string    `json:"complement"`
	Neighborhood    string    `json:"neighborhood"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	ZipCode         string    `json:"zip_code"`
	AddressableID   uuid.UUID `json:"addressable_id"`
	AddressableType string    `json:"addressable_type"`
	CreatedAt_3     time.Time `json:"created_at_3"`
	UpdatedAt_3     time.Time `json:"updated_at_3"`
}

type OwnerStore interface {
	CreateOwner(OwnerPayload) (uuid.UUID, error)
	GetOwners() ([]*Owner, error)
	GetOwner(id uuid.UUID) (*Owner, error)
}
