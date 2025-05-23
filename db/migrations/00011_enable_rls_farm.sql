-- +goose Up
-- +goose StatementBegin

ALTER TABLE farm ENABLE ROW LEVEL SECURITY;

CREATE POLICY farm_isolation_policy
  ON farm
  USING (
    current_setting('app.current_user_id')::INTEGER = user_id
  );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
