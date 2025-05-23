-- name: CreateCurrencyBalance :one
INSERT INTO "currency_balance" (user_id, currency_type_id, balance)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCurrencyBalance :one
SELECT *
FROM "currency_balance"
WHERE id = $1;

-- TODO: Add FOR UPDATE query to lock and prevent race conditoins, when needed

-- name: UpdateCurrencyBalance :one
UPDATE "currency_balance"
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCurrencyBalance :exec
DELETE FROM "currency_balance"
WHERE id = $1;

-- name: UpsertUserCurrencyBalance :one
INSERT INTO "currency_balance" (user_id, currency_type_id, balance)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, currency_type_id) DO UPDATE
SET balance = EXCLUDED.balance
RETURNING *;

-- name: GetUserCurrencyBalanceByType :one
SELECT *
FROM "currency_balance"
WHERE user_id = $1 AND currency_type_id = $2;

-- name: ListUserCurrencyBalances :many
SELECT *
FROM "currency_balance"
WHERE user_id = $1;

-- name: AdjustUserCurrencyBalanceByType :one
UPDATE "currency_balance"
SET balance = balance + sqlc.arg(amount)
WHERE user_id = sqlc.arg(user_id) AND currency_type_id = sqlc.arg(currency_type_id)
RETURNING *;