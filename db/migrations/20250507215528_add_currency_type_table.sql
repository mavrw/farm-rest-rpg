-- +goose Up
-- +goose StatementBegin
CREATE TABLE "currency_type" (
    "id"                    SERIAL  PRIMARY KEY,
    "name"                  TEXT    NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "currency_type";
-- +goose StatementEnd
