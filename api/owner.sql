-- name: CreateOwner :exec
INSERT INTO owners (ID, people_type, is_active, bucket_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetOwner :one
SELECT *
FROM owners
JOIN people ON owners.id = people.personable_id AND people.personable_type = 'owner'
JOIN addresses ON owners.id = addresses.addressable_id AND addresses.addressable_type = 'owner'
WHERE owners.id = $1;

-- name: GetOwners :many
SELECT *
FROM owners
JOIN people ON owners.id = people.personable_id AND people.personable_type = 'owner'
JOIN addresses ON owners.id = addresses.addressable_id AND addresses.addressable_type = 'owner';

-- name: UpdateOwner :exec
UPDATE owners SET people_type = $2, is_active = $3, bucket_id = $4, updated_at = $5 WHERE id = $1;

-- name: DisableOwner :exec
UPDATE owners SET is_active = $2, updated_at = $3 WHERE owners.id = $1;
