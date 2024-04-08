-- name: CreateAttendance :exec
INSERT INTO attendances ( ID, date_service, start_service, end_service, status, reminder, owner_id, type_service_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: GetAttendance :one
SELECT *
FROM attendances
WHERE attendances.id = $1;

-- name: GetAttendances :many
SELECT *
FROM attendances;

-- name: GetAttendanceDateService :one
SELECT *
FROM attendances
WHERE attendances.date_service = $1;

-- name: UpdateAttendance :exec
UPDATE attendances SET date_service = $2, start_service = $3, end_service = $4, status = $5, reminder = $6, type_service_id = $7, updated_at = $8 WHERE attendances.id = $1;

-- name: DeleteAttendance :exec
DELETE FROM attendances
WHERE attendances.id = $1
AND NOT EXISTS (
    SELECT 1 FROM owners WHERE owner_id = $1
)
AND NOT EXISTS (
    SELECT 1 FROM type_services WHERE type_service_id = $1
);
