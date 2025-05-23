-- +goose Up
-- +goose StatementBegin
CREATE TABLE "market_catalog" (
    item_id     INTEGER PRIMARY KEY REFERENCES "item"(id),
    buy_price   INTEGER,
    sell_price  INTEGER

    -- NULL prices signal item cannot be bought/sold
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "market_catalog";
-- +goose StatementEnd
