-- +goose Up
-- +goose StatementBegin
CREATE TABLE "inventory" (
    "id"    SERIAL  PRIMARY KEY,
    "user_id"   INTEGER REFERENCES "users"(id)   NOT NULL,
    "item_id"   INTEGER REFERENCES "item"(id)   NOT NULL,
    "quantity"  INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "inventory";
-- +goose StatementEnd
