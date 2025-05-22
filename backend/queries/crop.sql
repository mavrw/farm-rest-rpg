-- name: CreateCrop :one
INSERT INTO "crop" (id, name, growth_time_seconds, yield_amount)
VALUES ($1, $2, $3, $4)
ON CONFLICT (id) DO NOTHING
RETURNING *;

-- name: GetAllCrops :many
SELECT * 
FROM "crop"
ORDER BY id;

-- name: GetCropByID :one
SELECT *
FROM "crop"
WHERE id = $1;