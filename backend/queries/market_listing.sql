-- name: CreateMarketListing :one
INSERT INTO "market_listing" (item_id, buy_price, sell_price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetMarketListing :one
SELECT *
FROM "market_listing"
WHERE item_id = $1;

-- name: UpdateMarketListing :one
UPDATE "market_listing"
SET buy_price = $2,
    sell_price = $3
WHERE item_id = $1
RETURNING *;

-- name: DeleteMarketListing :exec
DELETE FROM "market_listing"
WHERE item_id = $1;

-- name: ListMarketListings :many
SELECT *
FROM "market_listing"
ORDER BY item_id;
-- TODO: Add pagination?

-- name: GetMarketListingBuyPrice :one
SELECT buy_price
FROM "market_listing"
WHERE item_id = $1;

-- name: GetMarketListingSellPrice :one
SELECT sell_price
FROM "market_listing"
WHERE item_id = $1;

-- name: MarketListingExists :one
SELECT EXISTS (
    SELECT 1
    FROM "market_listing"
    WHERE item_id = $1
) AS exists;