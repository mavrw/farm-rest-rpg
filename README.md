# Text-based Farming RPG

## Overview

This document outlines the technical design of a text-based farming RPG game designed for AWS deployment.  

The game allows users to manage their farms, plant and harvest crops, sell goods at the town market, and engage in other time-based activities.

---

## Goals & Non-Goals

### Goals

- Provide endpoints for account creation and login
- Provide endpoints for account management (e.g. password change, profile customizations)
- Provide endpoints for player inventory (view, add, destroy)
- Provide endpoints for farm management (e.g. buying plots, upgrades)
- Provide endpoints for planting and harvesting crops
- Provide endpoints for selling goods at the town market
- Secure API with JWT-based authentication
- Use PostgreSQL for persistent storage
- Implement simple web frontend interface for testing and interaction
- Setup Infrastructure as Code (IaC) and CI/CD for cloud deployment and rapid iteration

### Non-Goals

- Implement real-time chat or multi-player features
- Integrate payment services for microtransactions
- Enable direct player-to-player interactions or trading
- Add extended features such as animals, orchards, brewing, or crafting systems
- Support mobile platforms or native desktop clients
- Implement MFA sign-in authentication via HOTP

---

## Architecture Overview

This project consists of four major components:
  
- A RESTful API backend written in Go
- A relational database for persistent storage
- A web frontend that interacts with the backend via JSON APIs
- Cloud infrastructure for deployment

### Backend

The backend API will be written in Go using the [Gin] framework for routing and middleware. SQL queries and type-safe data access will be handled via sqlc, which generates Go code from raw SQL queries. The API will expose RESTful endpoints for interacting with the game logic and persistent data.

### Database

[PostgreSQL] will serve as the primary relational database for storing user accounts, farms, inventories, market transactions, and time-based game events. Schema versioning will be managed via `migrate`, `goose`, or similar tooling (to be finalized).

### Frontend

The initial frontend will be a minimal [Vue3] application, primarily used for testing API endpoints and simulating player interaction. Future iterations may enhance UI/UX or incorporate other frontend frameworks.

### AWS Infrastructure

Infrastructure will be defined and deployed using [Terraform] to ensure repeatability and version control. The architecture will initially consist of:

- An EC2 instance or container service (ECS/Fargate) to host the API
- RDS for PostgreSQL
- S3 for static asset hosting
- IAM roles and security groups for access control

### Architecture Diagram

```mermaid
graph TD
  subgraph Frontend
    A[Vue 3 App]
  end

  subgraph Backend
    B[Go + Gin API]
    B --> C[sqlc]
  end

  subgraph Database
    D[(PostgreSQL DB)]
  end

  subgraph AWS Infrastructure
    E[Terraform]
    F[EC2 or ECS/Fargate]
    G[RDS - PostgreSQL]
    H[S3 - optional]
    I[IAM Roles/Security Groups]
  end

  A -->|HTTP - JSON REST| B
  B -->|SQL Queries| D
  E --> H
  E --> F
  E --> G
  E --> I
  F --> B
  G --> D
  A --> H

```

---

## Data Models/Schema

### Tables Overview

```plaintext
users
- id (PK)
- username
- email
- password_hash
- created_at

farms
- id (PK)
- user_id (FK to users)
- name
- created_at

plots
- id (PK)
- farm_id (FK to farms)
- x
- y
- soil_type
- crop_id (FK to crops, nullable)
- planted_at

crops
- id (PK)
- name
- growth_time
- yield_amount
- season

animals
- id (PK)
- farm_id (FK to farms)
- type
- age
- health
- hunger
- last_fed_at

items
- id (PK)
- name
- type (tool, seed, feed, etc.)
- effect_json

inventory
- id (PK)
- user_id (FK to users)
- item_id (FK to items)
- quantity

```

> Future Expansions: `quests`, `weather`, `npc`, `player_exchange`, etc.

---

## API/Interface Design

### Public

- `POST /api/v1/register`  
- `POST /api/v1/login`  

### Protected (JWT)

- `GET  /api/v1/farm`  
- `POST /api/v1/plot/plant`  
- `POST /api/v1/plot/harvest`  
- `POST /api/v1/animal/feed`  
- `GET  /api/v1/inventory`  
- `POST /api/v1/market/buy`  
- `POST /api/v1/market/sell`  

---

## Design Details

### Time-based Crop Growth

> *What is the tick/loop model? How are time-based events handled?*

### Inventory Transaction Handling

> *How is atomicity handled? How are inventory rollbacks handled? How are potential buy/sell race conditions handled?*

### Data Access Patterns

> *Is RESTful CRUD sufficient, or should CQRS be explored as well?*

---

## Security & Performance

### Authentication/Authorization

> *What is the JWT expiration and refresh strategy?*

### Data validation/sanitization

> *How is SQL Injection prevented?*

### Performance Expectations or Bottlenecks

### Rate Limits

> *How is rate limiting implemented to prevent against DDOS or Brute Force attacks?*

### Caching

> *Does caching need to be implemented to buffer database operations?*

---

## Testing Plan

### Unit Tests

### Integration Tests

### Testing Tools/Frameworks

- `Postman`

---

## Deployment & Operations

### Environments

- `dev`
- `staging`
- `prod`

### Deployment Process

> *How will deployments take place?*

### Logging, Monitoring, and Error Tracking

> *How will loggins, monitoring, and error tracking be implemented?*

### Rollback and Recovery Plans

> *In the event of a deployment failure or other catastrophy, what is the rollback and recovery plan?*

---

## Risks & Mitigations

> *What potential technical or process risks can be identified? How can they be mitigated?*

---
