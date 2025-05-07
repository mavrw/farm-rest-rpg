.PHONY: help up down migrate build-backend build-frontend test logs

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	  awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

up:               ## Build & start all services (db, migrate, backend, frontend, nginx)
	docker compose up --build -d

down:             ## Tear down all services
	docker compose down

migrate:
	docker compose up migrate

build-backend:    ## Build only the backend image
	docker compose build backend

build-frontend:   ## Build only the frontend image
	docker compose build frontend

test:             ## Run backend tests (requires db up)
	docker compose up -d db
	docker compose exec backend go test ./...

logs:             ## Tail logs for all services
	docker compose logs -f
