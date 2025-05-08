-- name: CreateUser :exec
INSERT INTO "user" (username, email, password_hash)
VALUES ($1, $2, $3);

-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, created_at, updated_at
FROM "user"
WHERE email = $1;