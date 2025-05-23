-- name: CreateCropDefinition :one
INSERT INTO "crop" (
    id, 
    name, 
    growth_time_seconds, 
    seed_item_id, 
    yield_item_id, 
    yield_amount
)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id) DO NOTHING
RETURNING *;

-- name: ListCropDefinitions :many
SELECT * 
FROM "crop"
ORDER BY id;

-- name: GetCropDefinition :one
SELECT *
FROM "crop"
WHERE id = $1;

-- name: GetCropDefinitionByName :one
SELECT *
FROM "crop"
WHERE name = $1;
