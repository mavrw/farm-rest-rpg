-- name: CreateFarm :one
INSERT INTO "farm" (user_id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetFarmByID :one
SELECT *
FROM "farm"
WHERE id = $1;

-- name: GetFarmByUserID :one
SELECT *
FROM "farm"
WHERE user_id = $1;