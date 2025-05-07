-- +goose Up
-- +goose StatementBegin
CREATE TABLE "crop" (
    "id"            SERIAL  PRIMARY KEY,
    "name"          TEXT    NOT NULL,
    "growth_time"   INTEGER NOT NULL,
    "yield_amount"  INTEGER NOT NULL,
    "season"        TEXT    DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "crop";
-- +goose StatementEnd
