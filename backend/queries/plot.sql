-- name: CreatePlot :one
INSERT INTO "plot" (farm_id)
VALUES ($1)
RETURNING *;

-- name: GetPlotsByFarmID :many
SELECT *
FROM "plot" 
WHERE farm_id = $1;

-- name: GetPlotByID :one
SELECT * 
FROM "plot" 
WHERE id = $1;

-- name: SowPlotByID :one
UPDATE "plot"
SET crop_id = $2,
    planted_at = $3,
    harvest_at = $4
WHERE id = $1
RETURNING *;

-- name: HarvestPlotByID :one
UPDATE "plot"
SET crop_id = NULL,
    planted_at = NULL,
    harvest_at = NULL
WHERE id = $1
RETURNING *;