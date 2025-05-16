-- name: CreateUser :exec
INSERT INTO "users" (username, email, password_hash)
VALUES ($1, $2, $3);

-- name: GetUserByID :one
SELECT id, username, email, password_hash, created_at, updated_at
FROM "users"
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, created_at, updated_at
FROM "users"
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT id, username, email, password_hash, created_at, updated_at
FROM "users" 
WHERE username = $1;