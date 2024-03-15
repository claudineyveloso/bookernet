// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: owner.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createOwner = `-- name: CreateOwner :exec
INSERT INTO owners (ID, people_type, is_active, bucket_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateOwnerParams struct {
	ID         uuid.UUID `json:"id"`
	PeopleType string    `json:"people_type"`
	IsActive   bool      `json:"is_active"`
	BucketID   uuid.UUID `json:"bucket_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (q *Queries) CreateOwner(ctx context.Context, arg CreateOwnerParams) error {
	_, err := q.db.ExecContext(ctx, createOwner,
		arg.ID,
		arg.PeopleType,
		arg.IsActive,
		arg.BucketID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const disableOwner = `-- name: DisableOwner :exec
UPDATE owners SET is_active = $2, updated_at = $3 WHERE owners.id = $1
`

type DisableOwnerParams struct {
	ID        uuid.UUID `json:"id"`
	IsActive  bool      `json:"is_active"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) DisableOwner(ctx context.Context, arg DisableOwnerParams) error {
	_, err := q.db.ExecContext(ctx, disableOwner, arg.ID, arg.IsActive, arg.UpdatedAt)
	return err
}

const getOwner = `-- name: GetOwner :one
SELECT owners.id, people_type, is_active, bucket_id, owners.created_at, owners.updated_at, people.id, first_name, last_name, email, phone, cell_phone, personable_id, personable_type, people.created_at, people.updated_at, addresses.id, public_place, complement, neighborhood, city, state, zip_code, addressable_id, addressable_type, addresses.created_at, addresses.updated_at
FROM owners
JOIN people ON owners.id = people.personable_id AND people.personable_type = 'owner'
JOIN addresses ON owners.id = addresses.addressable_id AND addresses.addressable_type = 'owner'
WHERE owners.id = $1
`

type GetOwnerRow struct {
	ID              uuid.UUID      `json:"id"`
	PeopleType      string         `json:"people_type"`
	IsActive        bool           `json:"is_active"`
	BucketID        uuid.UUID      `json:"bucket_id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	ID_2            uuid.UUID      `json:"id_2"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Email           string         `json:"email"`
	Phone           sql.NullString `json:"phone"`
	CellPhone       string         `json:"cell_phone"`
	PersonableID    uuid.UUID      `json:"personable_id"`
	PersonableType  string         `json:"personable_type"`
	CreatedAt_2     time.Time      `json:"created_at_2"`
	UpdatedAt_2     time.Time      `json:"updated_at_2"`
	ID_3            uuid.UUID      `json:"id_3"`
	PublicPlace     sql.NullString `json:"public_place"`
	Complement      sql.NullString `json:"complement"`
	Neighborhood    sql.NullString `json:"neighborhood"`
	City            sql.NullString `json:"city"`
	State           sql.NullString `json:"state"`
	ZipCode         sql.NullString `json:"zip_code"`
	AddressableID   uuid.UUID      `json:"addressable_id"`
	AddressableType string         `json:"addressable_type"`
	CreatedAt_3     time.Time      `json:"created_at_3"`
	UpdatedAt_3     time.Time      `json:"updated_at_3"`
}

func (q *Queries) GetOwner(ctx context.Context, id uuid.UUID) (GetOwnerRow, error) {
	row := q.db.QueryRowContext(ctx, getOwner, id)
	var i GetOwnerRow
	err := row.Scan(
		&i.ID,
		&i.PeopleType,
		&i.IsActive,
		&i.BucketID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ID_2,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.CellPhone,
		&i.PersonableID,
		&i.PersonableType,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.ID_3,
		&i.PublicPlace,
		&i.Complement,
		&i.Neighborhood,
		&i.City,
		&i.State,
		&i.ZipCode,
		&i.AddressableID,
		&i.AddressableType,
		&i.CreatedAt_3,
		&i.UpdatedAt_3,
	)
	return i, err
}

const getOwners = `-- name: GetOwners :many
SELECT owners.id, people_type, is_active, bucket_id, owners.created_at, owners.updated_at, people.id, first_name, last_name, email, phone, cell_phone, personable_id, personable_type, people.created_at, people.updated_at, addresses.id, public_place, complement, neighborhood, city, state, zip_code, addressable_id, addressable_type, addresses.created_at, addresses.updated_at
FROM owners
JOIN people ON owners.id = people.personable_id AND people.personable_type = 'owner'
JOIN addresses ON owners.id = addresses.addressable_id AND addresses.addressable_type = 'owner'
`

type GetOwnersRow struct {
	ID              uuid.UUID      `json:"id"`
	PeopleType      string         `json:"people_type"`
	IsActive        bool           `json:"is_active"`
	BucketID        uuid.UUID      `json:"bucket_id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	ID_2            uuid.UUID      `json:"id_2"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Email           string         `json:"email"`
	Phone           sql.NullString `json:"phone"`
	CellPhone       string         `json:"cell_phone"`
	PersonableID    uuid.UUID      `json:"personable_id"`
	PersonableType  string         `json:"personable_type"`
	CreatedAt_2     time.Time      `json:"created_at_2"`
	UpdatedAt_2     time.Time      `json:"updated_at_2"`
	ID_3            uuid.UUID      `json:"id_3"`
	PublicPlace     sql.NullString `json:"public_place"`
	Complement      sql.NullString `json:"complement"`
	Neighborhood    sql.NullString `json:"neighborhood"`
	City            sql.NullString `json:"city"`
	State           sql.NullString `json:"state"`
	ZipCode         sql.NullString `json:"zip_code"`
	AddressableID   uuid.UUID      `json:"addressable_id"`
	AddressableType string         `json:"addressable_type"`
	CreatedAt_3     time.Time      `json:"created_at_3"`
	UpdatedAt_3     time.Time      `json:"updated_at_3"`
}

func (q *Queries) GetOwners(ctx context.Context) ([]GetOwnersRow, error) {
	rows, err := q.db.QueryContext(ctx, getOwners)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetOwnersRow
	for rows.Next() {
		var i GetOwnersRow
		if err := rows.Scan(
			&i.ID,
			&i.PeopleType,
			&i.IsActive,
			&i.BucketID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ID_2,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.CellPhone,
			&i.PersonableID,
			&i.PersonableType,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.ID_3,
			&i.PublicPlace,
			&i.Complement,
			&i.Neighborhood,
			&i.City,
			&i.State,
			&i.ZipCode,
			&i.AddressableID,
			&i.AddressableType,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOwner = `-- name: UpdateOwner :exec
UPDATE owners SET people_type = $2, is_active = $3, bucket_id = $4, updated_at = $5 WHERE id = $1
`

type UpdateOwnerParams struct {
	ID         uuid.UUID `json:"id"`
	PeopleType string    `json:"people_type"`
	IsActive   bool      `json:"is_active"`
	BucketID   uuid.UUID `json:"bucket_id"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (q *Queries) UpdateOwner(ctx context.Context, arg UpdateOwnerParams) error {
	_, err := q.db.ExecContext(ctx, updateOwner,
		arg.ID,
		arg.PeopleType,
		arg.IsActive,
		arg.BucketID,
		arg.UpdatedAt,
	)
	return err
}
