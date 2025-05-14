-- name: CreateCrop :exec
INSERT INTO "crop" (name, growth_time_seconds, yield_amount)
VALUES ($1, $2, $3)
ON CONFLICT (id) DO NOTHING;

-- name: GetAllCrops :many
SELECT * 
FROM "crop"
ORDER BY id;