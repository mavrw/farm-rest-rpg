-- +goose Up
-- +goose StatementBegin
CREATE TABLE "plot" (
    "id"            SERIAL  PRIMARY KEY,
    "farm_id"       INTEGER REFERENCES "farm"(id) NOT NULL,
    "x"             INTEGER NOT NULL,
    "y"             INTEGER NOT NULL,
    "soil_type"     TEXT    DEFAULT NULL,
    "crop_id"       INTEGER REFERENCES "crop"(id) DEFAULT NULL,
    "planted_at"    TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "plot";
-- +goose StatementEnd
