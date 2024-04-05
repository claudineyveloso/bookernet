-- name: CreateTypeService :exec
INSERT INTO type_services ( ID, name, duration, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5);

-- name: GetTypeService :one
SELECT *
FROM type_services
WHERE type_services.id = $1;

-- name: GetTypeServices :many
SELECT *
FROM type_services;

-- name: UpdateTypeService :exec
UPDATE type_services SET name = $2, duration = $3, updated_at = $4 WHERE type_services.id = $1;

-- name: DeleteTypeService :exec
DELETE FROM type_services
WHERE type_services.id = $1
AND NOT EXISTS (
    SELECT 1 FROM attendances WHERE type_service_id = $1
);
