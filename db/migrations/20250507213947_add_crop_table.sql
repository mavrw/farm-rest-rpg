-- +goose Up
-- +goose StatementBegin
CREATE TABLE "crop" (
    "id"                    SERIAL  PRIMARY KEY,
    "name"                  TEXT    NOT NULL,
    "growth_time_seconds"   INTEGER NOT NULL,
    "yield_amount"          INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "crop";
-- +goose StatementEnd
