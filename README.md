# Push Notification Service (PNS)

Backend service that exposes a mobile-facing device registration API and a simple web admin UI for managing users and promotions. Built with Go + Gin, PostgreSQL for storage, and ships with Docker Compose for local/dev deployment.

## Features
- Register device tokens via REST (`/api/v1/devices/register`), secured by a shared `SERVER_KEY` header.
- Ensures only one active FCM token per device; older tokens are automatically marked inactive.
- Web admin UI with session auth for:
  - User CRUD (manage who can log into the admin).
  - Promotion CRUD (title, image URL, description, start/end period, status).
  - Activate a promotion, which triggers push notifications to all active device tokens.
- DB migrations run automatically on start; default admin seeding via env vars.

## Stack
- Go 1.22, Gin, pgx, bcrypt.
- PostgreSQL 15.
- Docker + Docker Compose for repeatable runtime.

## Running with Docker (recommended)
```bash
cd /Users/ziaulhaq/dev/project/INAINA/BE/pns
docker compose up --build
```
The app listens on `http://localhost:8080`.

Environment (tweak in `docker-compose.yml` or export when running without Compose):
- `APP_PORT` (default `8080`)
- `DATABASE_URL` (default points to compose `db`)
- `SERVER_KEY` (required; header to authorize device API calls)
- `SESSION_SECRET` (required; cookie signing secret)
- `ADMIN_EMAIL` / `ADMIN_PASSWORD` (optional; creates a bootstrap admin if not present)

## Manual run (needs Go installed)
```bash
export DATABASE_URL=postgres://pns:pns@localhost:5432/pns?sslmode=disable
export SERVER_KEY=dev-server-key
export SESSION_SECRET=change-me
export ADMIN_EMAIL=admin@example.com
export ADMIN_PASSWORD=admin123
go run ./cmd/server
```

## REST API
- `POST /api/v1/devices/register`
  - Headers: `X-Server-Key: <SERVER_KEY>`
  - Body: `{"device_id":"<device>","fcm_device_token":"<token>"}`
  - Response: stored token record; any other token for the same `device_id` is marked inactive.

## Admin UI
- Login: `http://localhost:8080/login` (use the bootstrap admin credentials).
- Users: list/create/edit/delete.
- Promotions: list/create/edit; activating sends push to active device tokens (using provided `SERVER_KEY` as the FCM key).

## Database schema (auto-migrated)
- `users`: id, email (unique), password_hash, created_at.
- `device_tokens`: id, device_id, fcm_device_token, status (active|inactive), created_at; enforces one active per device.
- `promotions`: id, title, image_url, description, start_date, end_date, status (created|active|inactive), created_at.

## Notes
- The FCM sender is a minimal HTTP client; replace the API URL/key with your production values.
- All timestamps are UTC; password hashes use bcrypt.

## Development tips
- `go mod tidy` to sync dependencies (requires internet).
- Run `docker compose down -v` to reset the local database.

