-- name: CreatePerson :exec
INSERT INTO people ( ID, first_name, last_name, email, phone, cell_phone, personable_id, personable_type, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: UpdatePerson :exec
UPDATE people SET first_name = $2, last_name = $3, updated_at = $4 WHERE id = $1;
