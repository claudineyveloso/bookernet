-- name: CreateCustomer :exec
INSERT INTO customers (ID, birthday, created_at, updated_at)
VALUES ($1, $2, $3, $4);

-- name: GetCustomer :one
SELECT *
FROM customers
JOIN people ON customers.id = people.personable_id AND people.personable_type = 'customer'
JOIN addresses ON customers.id = addresses.addressable_id AND addresses.addressable_type = 'customer'
WHERE customers.id = $1;

-- name: GetCustomers :many
SELECT *
FROM customers
JOIN people ON customers.id = people.personable_id AND people.personable_type = 'customer'
JOIN addresses ON customers.id = addresses.addressable_id AND addresses.addressable_type = 'customer';

-- name: UpdateCustomer :exec
UPDATE customers SET birthday = $2, updated_at = $3 WHERE id = $1;
