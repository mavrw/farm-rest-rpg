# Text-based Farming RPG

A monorepo for a text-based farming RPG game, featuring:

- **Go + Gin** backend API  
- **PostgreSQL** database with migrations & `sqlc`-generated models  
- **Vue 3 + TypeScript** frontend for testing/admin  
- **Nginx** reverse-proxy for unified frontend/backend hosting  
- **Docker Compose** for local development  
- **Terraform** stubs for AWS IaC  

> 📄 The detailed design document has been renamed to [`DESIGN-DOC.md`](./DESIGN-DOC.md).

---

## 🚀 Quick Start

1. **Copy & populate your credentials**  

    ```bash
    cp devops/.env.example devops/.env
    # edit devops/.env → set POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB
    ```

2. **Build & start everything**

    ```bash
    make up
    ```

3. **Visit the app**

   - Frontend + API proxy:  [http://localhost/](http://localhost/)
   - Go backend directly:   [http://localhost/api/hello](http://localhost/api/hello)
   - Adminer (DB GUI):      [http://localhost:8080](http://localhost:8080)

4. **Run tests**

   ```bash
   make test
   ```

5. **Stop & clean up**

   ```bash
   make down
   ```

---

## 📁 Repo Layout

```plaintext
text-farming-rpg/
├── backend/             # Go + Gin API
│   ├── cmd/api          # main.go entrypoint
│   ├── internal/        # business logic, auth, farm, inventory, db
│   ├── api/             # HTTP handlers & routes
│   ├── migrations/      # SQL migration files
│   ├── sqlc.yaml        # sqlc config
│   └── Dockerfile
│
├── frontend/            # Vue 3 + TypeScript app
│   ├── public/
│   ├── src/             # components, views, services
│   ├── vite.config.ts
│   └── Dockerfile
│
├── db/                  # Database tooling
│   ├── migrations/      # migrate/goose files
│   ├── seed.sql
│   └── fixtures/
│
├── devops/              # CI, Terraform, scripts, config
│   ├── terraform/
│   ├── scripts/
│   └── .env             # git-ignored credentials
│
├── nginx/               # Nginx reverse-proxy config
│   └── nginx.conf
│
├── docker-compose.yml   # local development stack
├── Makefile             # helper commands: up, down, build, test, logs
├── design-doc.md        # Technical design document
└── README.md            # (this file)
```

---

## 🛠 Commands

From the project root:

| Command                           | Description                                                           |
| --------------------------------- | --------------------------------------------------------------------- |
| `make up`                         | Build images & start all services (db, backend, frontend, nginx)      |
| `make down`                       | Stop & remove all containers                                          |
| `make build-backend`              | Build only the backend Docker image                                   |
| `make build-frontend`             | Build only the frontend Docker image                                  |
| `make test`                       | Run backend tests (requires `db` up)                                  |
| `make logs`                       | Tail logs for all services                                            |
| `docker compose exec backend ...` | Run arbitrary commands inside the backend container (e.g. migrations) |

---

## 🔧 Development Workflow

- **Migrations**

  - Write new SQL files in `db/migrations/`
  - Apply locally with `make up` (runs migrations on startup)
- **SQLC Models**

  - Define SQL queries in `backend/internal/db/queries/`
  - Regenerate with `cd backend && sqlc generate`
- **Frontend**

  - Develop in `<repo>/frontend/src/`
  - Dev server available at [http://localhost:3000](http://localhost:3000) (proxied by Nginx)
- **Terraform**

  - Stub modules in `devops/terraform/` for AWS deployment

---

## 📝 Notes

- **Environment variables** live in `devops/.env` (not committed).
- **No volume mounts** overwrite built images—our Dockerfiles produce immutable artifacts.
- **Nginx** proxies `/` to the frontend and `/api/` to the backend for a single-origin setup.

---
