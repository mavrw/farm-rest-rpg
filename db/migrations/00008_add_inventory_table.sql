-- +goose Up
-- +goose StatementBegin
CREATE TABLE "inventory" (
    "id"    SERIAL  PRIMARY KEY,
    "user_id"   INTEGER REFERENCES "users"(id)   NOT NULL,
    "item_id"   INTEGER REFERENCES "item"(id)   NOT NULL,
    "quantity"  INTEGER NOT NULL,

    -- add prevent a user from having two entries for one item
    UNIQUE ("user_id", "item_id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "inventory";
-- +goose StatementEnd
