-- +goose Up
-- +goose StatementBegin
CREATE TYPE item_rarity AS ENUM (
    'common',
    'uncommon',
    'rare',
    'epic',
    'legendary'
);

CREATE TABLE "item" (
    "id"            SERIAL      PRIMARY KEY,
    "name"          TEXT        UNIQUE NOT NULL,
    "description"   TEXT        NOT NULL DEFAULT '',
    "rarity"        item_rarity NOT NULL, 
    "type"          INTEGER     REFERENCES "item_type"(id) NOT NULL,
    "effect_json"   JSON        DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "item";
DROP TYPE item_rarity;
-- +goose StatementEnd
