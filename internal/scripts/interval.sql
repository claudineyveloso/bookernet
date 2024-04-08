-- name: CreateInterval :exec
INSERT INTO intervals ( ID, owner_id, interval_minutes, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5);

-- name: GetInterval :one
SELECT *
FROM intervals
WHERE intervals.id = $1;

-- name: GetIntervals :many
SELECT *
FROM intervals;

-- name: UpdateInterval :exec
UPDATE intervals SET owner_id = $2, updated_at = $3 WHERE intervals.id = $1;

-- name: DeleteInterval :exec
DELETE FROM intervals
WHERE intervals.id = $1
AND NOT EXISTS (
    SELECT 1 FROM owners WHERE owner_id = $1
);
