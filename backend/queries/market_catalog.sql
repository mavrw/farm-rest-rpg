-- name: CreateMarketItem :one
INSERT INTO "market_catalog" (item_id, buy_price, sell_price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetMarketItem :one
SELECT *
FROM "market_catalog"
WHERE item_id = $1;

-- name: UpdateMarketItem :one
UPDATE "market_catalog"
SET buy_price = $2,
    sell_price = $3
WHERE item_id = $1
RETURNING *;

-- name: DeleteMarketItem :exec
DELETE FROM "market_catalog"
WHERE item_id = $1;

-- name: ListMarketItems :many
SELECT *
FROM "market_catalog"
ORDER BY item_id;
-- TODO: Add pagination?

-- name: GetMarketBuyPrice :one
SELECT buy_price
FROM "market_catalog"
WHERE item_id = $1;

-- name: GetMarketSellPrice :one
SELECT sell_price
FROM "market_catalog"
WHERE item_id = $1;

-- name: MarketItemExists :one
SELECT EXISTS (
    SELECT 1
    FROM "market_catalog"
    WHERE item_id = $1
) AS exists;