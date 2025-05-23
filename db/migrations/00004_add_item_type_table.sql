-- +goose Up
-- +goose StatementBegin
CREATE TABLE "item_type" (
    "id"            SERIAL      PRIMARY KEY,
    "name"          TEXT        UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "item_type";
-- +goose StatementEnd
