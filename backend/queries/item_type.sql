-- name: CreateItemType :one
INSERT INTO "item_type" (id, name)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: GetAllItemTypes :many
SELECT * 
FROM "item_type"
ORDER BY id;

-- name: GetItemTypeByID :one
SELECT *
FROM "item_type"
WHERE id = $1;