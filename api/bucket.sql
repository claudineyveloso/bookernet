-- name: CreateBucket :exec
INSERT INTO buckets ( ID, description, name, aws_access_key_id, aws_secret_access_key, aws_region, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetBucket :one
SELECT *
FROM buckets
WHERE buckets.id = $1;

-- name: GetBuckets :many
SELECT *
FROM buckets;

-- name: UpdateBucket :exec
UPDATE buckets SET description = $2, name = $3, aws_access_key_id = $4, aws_secret_access_key = $5, aws_region = $6, updated_at = $7 WHERE buckets.id = $1;
