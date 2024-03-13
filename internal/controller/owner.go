package controller

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/google/uuid"
)

type AddressRequest struct {
	ID              uuid.UUID      `json:"id"`
	PublicPlace     sql.NullString `json:"public_place"`
	Complement      sql.NullString `json:"complement"`
	Neighborhood    sql.NullString `json:"neighborhood"`
	City            sql.NullString `json:"city"`
	State           sql.NullString `json:"state"`
	ZipCode         sql.NullString `json:"zip_code"`
	AddressableID   uuid.UUID      `json:"addressable_id"`
	AddressableType string         `json:"addressable_type"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type PersonRequest struct {
	ID             uuid.UUID      `json:"id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Email          string         `json:"email"`
	Phone          sql.NullString `json:"phone"`
	CellPhone      string         `json:"cell_phone"`
	PersonableID   uuid.UUID      `json:"personable_id"`
	PersonableType string         `json:"personable_type"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type OwnerRequest struct {
	ID             uuid.UUID    `json:"id"`
	PeopleType     string       `json:"people_type"`
	IsActive       sql.NullBool `json:"is_active"`
	BucketID       uuid.UUID    `json:"bucket_id"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	PersonRequest  `json:"person"`
	AddressRequest `json:"address"`
}

func CreateOwnerController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {

}
