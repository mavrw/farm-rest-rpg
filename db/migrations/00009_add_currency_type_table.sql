-- +goose Up
-- +goose StatementBegin
CREATE TABLE "currency_type" (
    "id"    INTEGER     PRIMARY KEY CHECK (id > 0),
    "name"  TEXT        NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "currency_type";
-- +goose StatementEnd
