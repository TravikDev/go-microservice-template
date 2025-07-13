# Go Microservice Template

This repository provides a starting point for Go-based microservices using Temporal and NATS with a simple REST API.

## Project Structure

```
/user-service
├── cmd/
│   ├── api/          # Starts the HTTP REST server
│   ├── worker/       # Starts the Temporal worker
│   └── subscriber/   # Starts the NATS subscriber
├── api/
│   ├── handler/      # HTTP handlers
│   ├── middleware/   # HTTP middlewares
│   └── router/       # Routing setup
├── config/           # Configuration loading
├── controller/       # Orchestrates use cases
├── service/          # Business logic
├── workflow/         # Temporal workflows and activities
├── subscriber/       # NATS message handlers
├── repository/       # Database interfaces and implementations
├── model/            # Domain models
├── client/           # External service clients
├── internal/         # Private helpers
├── test/             # Integration tests
```

## Local Checks

Run `./scripts/check.sh` to format, vet and test the project.

```bash
./scripts/check.sh
```

## CI

GitHub Actions will execute the same checks on every push or pull request via `.github/workflows/checks.yml`.

## Sample API

- `GET /ping` - health check
- `GET /users` - list all users
- `GET /users/{id}/characters` - list characters for a user
## Installation

Ensure Go 1.24 or newer is installed. Download dependencies with:

```bash
make deps
```

## Environment Variables

- `NATS_URL` - NATS server URL
- `TEMPORAL_URL` - Temporal frontend address
- `DATABASE_URL` - PostgreSQL connection string

## Running the Services

Each component can be started individually:

```bash
make api        # launches the REST API on :8080
make worker     # runs the Temporal worker
make subscriber # starts the NATS subscriber
```
