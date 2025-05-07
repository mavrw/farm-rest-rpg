-- +goose Up
-- +goose StatementBegin
CREATE TABLE "item" (
    "id"            SERIAL  PRIMARY KEY,
    "name"          TEXT    NOT NULL,
    "tool"          TEXT    NOT NULL,
    "effect_json"   JSON    DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "item";
-- +goose StatementEnd
