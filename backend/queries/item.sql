-- name: CreateItemDefinition :one
INSERT INTO "item" (
    id, 
    name, 
    description,
    rarity,
    type, 
    effect_json
)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id) DO NOTHING
RETURNING *;

-- name: ListItemDefinitions :many
SELECT *
FROM "item"
ORDER BY id;

-- name: GetItemDefinition :one
SELECT *
FROM "item"
WHERE id = $1;

-- name: GetItemDefinitionByName :one
SELECT *
FROM "item"
WHERE name = $1;
