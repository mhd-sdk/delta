## Delta


### Architecture

- Go backend: `backend/`
  - HTTP API and WebSocket (Fiber)
  - gRPC slave discovery over UDP and log subscription
  - Alpaca integration (routing/market data)
  - Protobuf/gRPC (sources in `backend/proto`, generated clients in `backend/internal/pb`)
- React + Vite frontend: `frontend/`
- Deployment/K8s: `kube/`

---

## Prerequisites

- Go 1.23+
- Node.js 20+ and a package manager (yarn or npm)
- Docker and Docker Compose (optional, for containerized services)
- Protobuf Compiler `protoc` (to regenerate gRPC stubs if needed)
- Mage (Go task runner)
- API keys for Alpaca (for backend)

### Install tooling

```bash
# Protoc (macOS)
brew install protobuf

# Go plugins for protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Mage
go install github.com/magefile/mage@latest
```

Ensure `$GOPATH/bin` is on your `$PATH` so that `protoc-gen-go`, `protoc-gen-go-grpc`, and `mage` are available.

---

## Installation

```bash
git clone https://github.com/mhd-sdk/delta.git
cd Delta
```

### Backend configuration

The backend loads a `.env` file at the root of `backend/` (via `github.com/joho/godotenv`). Required variables:

```env
# backend/.env
API_KEY=YourAlpacaApiKey
SECRET_KEY=YourAlpacaSecretKey
MARKET_DATA_URL=https://data.alpaca.markets
ROUTING_URL=https://paper-api.alpaca.markets
```

Without these variables the backend will not start (`ErrMissingEnvVars`).

---

## Development

### Backend (Go)

```bash
cd backend
# Run directly
go run ./cmd/main.go

# Or build then run
go build -o main ./cmd/main.go
./main
```

By default, the HTTP API and WebSocket listen on `:3000`.

### Frontend (React/Vite)

```bash
cd frontend
yarn install      # or npm install
yarn dev          # starts Vite on http://localhost:5173
```

---

## Mage commands (backend)

From the `backend/` directory:

```bash
# List commands
mage -l

# Run a command with verbose logs
mage -v <target>
```

Available commands:

- `Dev`: runs the app with `go run`. Note: `magefile.go` references `cmd/delta/delta.go`, while this repo's actual entrypoint is `cmd/main.go`. If `Dev` fails due to a missing file, use `go run ./cmd/main.go` instead or update `magefile.go`.
- `Build`: builds the binary. Same note as above regarding the entrypoint path.
- `Proto`: generates gRPC stubs from `proto/state.proto`.

Examples:

```bash
cd backend
mage proto        # generates internal/pb from proto/state.proto
mage build        # builds the executable (if the entry path is aligned)
```

---

## API and WebSocket

Default backend base URL: `http://localhost:3000`

### Logs WebSocket

- Endpoint: `ws://localhost:3000/ws/logs`
- Purpose: subscribes to logs coming from gRPC slaves discovered over UDP `:9999`.

### Market — Bars (candlesticks)

- Endpoint: `POST /api/market-data/bars`
- JSON body:

```json
{
  "data": {
    "symbol": "AAPL",
    "start": "2024-01-01T09:30:00Z",
    "end": "2024-01-01T16:00:00Z",
    "timeframe": { "n": 1, "unit": "Minute" }
  }
}
```

Response: array of bars as returned by the Alpaca Market Data SDK.

Notes:
- The `symbol` field is the ticker.
- `end` can be omitted (the backend currently comments out `End`).
- `timeframe.unit` must match Alpaca-supported units (e.g., `Minute`, `Hour`, `Day`).

### Algorithms — Update state

- Endpoint: `POST /api/algorithms/state`
- JSON body:

```json
{ "id": "slave-id", "status": "start|stop|..." }
```

Effect: sends a gRPC `Control` command to the identified slave if it has been discovered.

---

## Docker

### Backend only

```bash
cd backend
docker build -t delta-backend .
docker run --rm --env-file ./.env -p 3000:3000 delta-backend
```

### Docker Compose
(Not ready yet)

The `docker-compose.yml` file provides a ready-to-use PostgreSQL service. The `backend` and `frontend` services are present but commented out. To run Postgres only:

```bash
docker-compose up -d postgres
```

To run the full stack via Compose, uncomment the `backend` and `frontend` sections, adjust environment variables, then:

```bash
docker-compose up --build
```

Note: the backend currently does not connect to a database — Postgres is provided as a base for future evolution.

---

## Protobuf/gRPC

Sources: `backend/proto/state.proto`

Generate stubs (Go):

```bash
cd backend
mage proto
# or directly
protoc --go_out=. --go-grpc_out=. proto/state.proto
```

Generated artifacts live in `backend/internal/pb`.

---

## Frontend scripts

In `frontend/package.json`:

- `dev`: start Vite
- `build`: TypeScript + Vite build
- `lint`: ESLint
- `test`: Jest
- `preview`: Vite preview
- `format` / `format:check`: Prettier
- `knip`: unused code/imports analyzer

---

## Quick access

- Frontend: `http://localhost:5173`
- Backend API & WS: `http://localhost:3000`

---
