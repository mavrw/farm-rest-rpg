-- +goose Up
-- +goose StatementBegin
CREATE TABLE "crop" (
    "id"                    SERIAL  PRIMARY KEY,
    "name"                  TEXT    UNIQUE NOT NULL,
    "growth_time_seconds"   INTEGER NOT NULL,
    "seed_item_id"          INTEGER REFERENCES "item"(id) NOT NULL,
    "yield_item_id"         INTEGER REFERENCES "item"(id) NOT NULL,
    "yield_amount"          INTEGER NOT NULL DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "crop";
-- +goose StatementEnd
