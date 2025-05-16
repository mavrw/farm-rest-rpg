-- +goose Up
-- +goose StatementBegin

CREATE TABLE "refresh_token" (
  id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id       INTEGER     REFERENCES "user"(id) ON DELETE CASCADE NOT NULL,
  token         TEXT        NOT NULL UNIQUE,
  expires_at    TIMESTAMPTZ NOT NULL,
  revoked       BOOLEAN     DEFAULT FALSE,
  created_at    TIMESTAMPTZ DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "refresh_token";
-- +goose StatementEnd
