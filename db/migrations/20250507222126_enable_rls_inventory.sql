-- +goose Up
-- +goose StatementBegin

ALTER TABLE inventory ENABLE ROW LEVEL SECURITY;

CREATE POLICY inventory_isolation_policy
  ON inventory
  USING (
    current_setting('app.current_user_id')::INTEGER = user_id
  );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
