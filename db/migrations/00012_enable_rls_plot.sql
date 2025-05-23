-- +goose Up
-- +goose StatementBegin

ALTER TABLE plot ENABLE ROW LEVEL SECURITY;

CREATE POLICY plot_isolation_policy
  ON plot
  USING (
    EXISTS (
      SELECT 1
      FROM farm f
      WHERE f.id = plot.farm_id
        AND f.user_id = current_setting('app.current_user_id')::INTEGER
    )
  );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
