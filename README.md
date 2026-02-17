# CoCDB

RESTful API serving Clash of Clans game data, built with [Go](https://go.dev/) and [chi](https://github.com/go-chi/chi).

## Prerequisites

- Go 1.25 or higher
- Git

## Quick Start

```bash
# Clone the repository
git clone https://github.com/flapjacck/CoCDB.git
cd CoCDB

# Install dependencies
go mod download

# Run the server
go run .
```

The API starts on `http://localhost:3000`

## API Endpoints

### General

| Method | Path           | Description                          |
|--------|----------------|--------------------------------------|
| GET    | `/`            | API information and available routes |
| GET    | `/health`      | Health check with uptime and version |
| GET    | `/favicon.ico` | Favicon (place file in `static/`)    |

### Buildings — `/api/v1/buildings`

| Method | Path                                  | Description                     |
|--------|---------------------------------------|---------------------------------|
| GET    | `/api/v1/buildings`                   | List all building categories    |
| GET    | `/api/v1/buildings/{category}`        | List buildings in a category    |
| GET    | `/api/v1/buildings/{category}/{name}` | Get a specific building's data  |

**Categories:** `army`, `defensive`, `resource`, `traps`

**Examples:**

```bash
curl http://localhost:3000/api/v1/buildings
curl http://localhost:3000/api/v1/buildings/defensive/cannon
```

### Troops — `/api/v1/troops`

| Method | Path                               | Description                  |
|--------|-------------------------------------|------------------------------|
| GET    | `/api/v1/troops`                   | List all troop categories    |
| GET    | `/api/v1/troops/{category}`        | List troops in a category    |
| GET    | `/api/v1/troops/{category}/{name}` | Get a specific troop's data  |

**Categories:** `elixir`, `dark_elixir`, `super`

**Examples:**

```bash
curl http://localhost:3000/api/v1/troops
curl http://localhost:3000/api/v1/troops/elixir
```

### Response Format

All responses use a consistent JSON envelope:

```json
{
  "status": "success",
  "data": { ... },
  "meta": {
    "version": "1.0.0",
    "cached": false
  }
}
```

Error responses:

```json
{
  "status": "error",
  "error": {
    "code": 404,
    "message": "building not found: nonexistent"
  }
}
```

## Configuration

All settings are controlled via environment variables. Copy `.env.example` to `.env` for reference.

| Variable        | Default       | Description                          |
|-----------------|---------------|--------------------------------------|
| `PORT`          | `3000`        | Server listen port                   |
| `ENVIRONMENT`   | `development` | `development` or `production`        |
| `LOG_LEVEL`     | `info`        | `debug`, `info`, `warn`, `error`     |
| `DATA_DIR`      | `data`        | Path to the JSON data directory      |
| `APP_VERSION`   | `1.0.0`       | Version string returned by the API   |
| `CACHE_TTL`     | `5m`          | Cache time-to-live (Go duration)     |
| `CORS_ORIGINS`  | `*`           | Comma-separated allowed CORS origins |
| `READ_TIMEOUT`  | `10s`         | HTTP read timeout                    |
| `WRITE_TIMEOUT` | `10s`         | HTTP write timeout                   |
| `IDLE_TIMEOUT`  | `120s`        | HTTP idle timeout                    |

## Attribution & Licensing

We use material from the Clash of Clans Wiki at Fandom to make this database. Please see their licensing terms below:

Community content on Fandom is available under the Creative Commons Attribution-ShareAlike license (CC BY‑SA) unless otherwise noted. See: <https://clashofclans.fandom.com/wiki/Clash_of_Clans_Wiki>

Some images or other media may have separate licenses; always check the media page on Fandom before reuse.

All JSON files include a `source`, `source_url`, and `source_license` field pointing to the original wiki article and license.
