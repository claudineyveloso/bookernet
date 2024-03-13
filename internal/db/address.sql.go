// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: address.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createAddress = `-- name: CreateAddress :exec
INSERT INTO addresses ( ID, public_place, complement, neighborhood, city, state, zip_code, addressable_id, addressable_type, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`

type CreateAddressParams struct {
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

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) error {
	_, err := q.db.ExecContext(ctx, createAddress,
		arg.ID,
		arg.PublicPlace,
		arg.Complement,
		arg.Neighborhood,
		arg.City,
		arg.State,
		arg.ZipCode,
		arg.AddressableID,
		arg.AddressableType,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const updateAddress = `-- name: UpdateAddress :exec
UPDATE addresses SET public_place = $2, complement = $3, neighborhood = $4, city = $5, state = $6, zip_code = $7, updated_at = $8 WHERE id = $1
`

type UpdateAddressParams struct {
	ID           uuid.UUID      `json:"id"`
	PublicPlace  sql.NullString `json:"public_place"`
	Complement   sql.NullString `json:"complement"`
	Neighborhood sql.NullString `json:"neighborhood"`
	City         sql.NullString `json:"city"`
	State        sql.NullString `json:"state"`
	ZipCode      sql.NullString `json:"zip_code"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (q *Queries) UpdateAddress(ctx context.Context, arg UpdateAddressParams) error {
	_, err := q.db.ExecContext(ctx, updateAddress,
		arg.ID,
		arg.PublicPlace,
		arg.Complement,
		arg.Neighborhood,
		arg.City,
		arg.State,
		arg.ZipCode,
		arg.UpdatedAt,
	)
	return err
}
