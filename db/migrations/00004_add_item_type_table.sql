-- +goose Up
-- +goose StatementBegin
CREATE TABLE "item_type" (
    "id"    INTEGER     PRIMARY KEY CHECK (id > 0),
    "name"  TEXT        UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "item_type";
-- +goose StatementEnd
