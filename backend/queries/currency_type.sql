-- name: CreateCurrencyType :one
INSERT INTO "currency_type" (id, name)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: GetAllCurrencyTypes :many
SELECT * 
FROM "currency_type"
ORDER BY id;

-- name: GetCurrencyTypeByID :one
SELECT *
FROM "currency_type"
WHERE id = $1;