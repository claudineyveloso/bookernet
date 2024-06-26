// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: type_service.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTypeService = `-- name: CreateTypeService :exec
INSERT INTO type_services ( ID, name, duration, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
`

type CreateTypeServiceParams struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Duration  int32     `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) CreateTypeService(ctx context.Context, arg CreateTypeServiceParams) error {
	_, err := q.db.ExecContext(ctx, createTypeService,
		arg.ID,
		arg.Name,
		arg.Duration,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const deleteTypeService = `-- name: DeleteTypeService :exec
DELETE FROM type_services
WHERE type_services.id = $1
AND NOT EXISTS (
    SELECT 1 FROM attendances WHERE type_service_id = $1
)
`

func (q *Queries) DeleteTypeService(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTypeService, id)
	return err
}

const getTypeService = `-- name: GetTypeService :one
SELECT id, name, duration, created_at, updated_at
FROM type_services
WHERE type_services.id = $1
`

func (q *Queries) GetTypeService(ctx context.Context, id uuid.UUID) (TypeService, error) {
	row := q.db.QueryRowContext(ctx, getTypeService, id)
	var i TypeService
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Duration,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTypeServiceByName = `-- name: GetTypeServiceByName :one
SELECT id, name, duration, created_at, updated_at
FROM type_services
WHERE type_services.name = $1
`

func (q *Queries) GetTypeServiceByName(ctx context.Context, name string) (TypeService, error) {
	row := q.db.QueryRowContext(ctx, getTypeServiceByName, name)
	var i TypeService
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Duration,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTypeServices = `-- name: GetTypeServices :many
SELECT id, name, duration, created_at, updated_at
FROM type_services
`

func (q *Queries) GetTypeServices(ctx context.Context) ([]TypeService, error) {
	rows, err := q.db.QueryContext(ctx, getTypeServices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TypeService
	for rows.Next() {
		var i TypeService
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Duration,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateTypeService = `-- name: UpdateTypeService :exec
UPDATE type_services SET name = $2, duration = $3, updated_at = $4 WHERE type_services.id = $1
`

type UpdateTypeServiceParams struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Duration  int32     `json:"duration"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) UpdateTypeService(ctx context.Context, arg UpdateTypeServiceParams) error {
	_, err := q.db.ExecContext(ctx, updateTypeService,
		arg.ID,
		arg.Name,
		arg.Duration,
		arg.UpdatedAt,
	)
	return err
}
