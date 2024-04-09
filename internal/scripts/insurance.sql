-- name: CreateInsurance :exec
INSERT INTO insurances ( ID, name, period, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5);

-- name: GetInsurance :one
SELECT *
FROM insurances
WHERE insurances.id = $1;

-- name: GetInsurances :many
SELECT *
FROM insurances;

-- name: GetInsuranceByName :one
SELECT *
FROM insurances
WHERE insurances.name = $1;

-- name: UpdateInsurance :exec
UPDATE insurances SET name = $2, period = $3, updated_at = $4 WHERE insurances.id = $1;

-- name: DeleteInsurance :exec
DELETE FROM insurances
WHERE insurances.id = $1;
