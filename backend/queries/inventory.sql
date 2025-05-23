-- name: AddInventoryItem :one
INSERT INTO "inventory" (user_id, item_id, quantity)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, item_id) DO NOTHING
RETURNING *;

-- name: GetInventoryItem :one
SELECT *
FROM "inventory"
WHERE user_id = $1 AND item_id = $2;

-- name: UpdateInventoryItem :one
UPDATE "inventory"
SET user_id = $2,
    item_id = $3,
    quantity = $4
WHERE id = $1
RETURNING *;

-- name: UpsertInventoryItem :one
INSERT INTO inventory (user_id, item_id, quantity)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, item_id) DO UPDATE
SET quantity = inventory.quantity + EXCLUDED.quantity
RETURNING *;

-- name: RemoveInventoryItem :exec
DELETE FROM "inventory"
WHERE item_id = $2 AND user_id = $1;

-- name: ListInventoryItems :many
SELECT *
FROM "inventory"
WHERE user_id = $1;

-- name: SetInventoryItemQuantity :one
INSERT INTO "inventory" (user_id, item_id, quantity)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, item_id) DO UPDATE
SET quantity = EXCLUDED.quantity
RETURNING *;

-- name: HasItemQuantity :one
SELECT quantity >= $3 AS has_enough
FROM "inventory"
WHERE user_id = $1 AND item_id = $2;
