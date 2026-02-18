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
| GET    | `/health`      | Health check with uptime info      |
| GET    | `/favicon.ico` | Favicon (place file in `static/`)    |

### Buildings — `/api/buildings`

| Method | Path                                 | Description                     |
|--------|--------------------------------------|---------------------------------|
| GET    | `/api/buildings`                     | List all building categories    |
| GET    | `/api/buildings/{category}`          | List buildings in a category    |
| GET    | `/api/buildings/{category}/{name}`   | Get a specific building's data  |

**Categories:** `army`, `defensive`, `resource`, `traps`

**Examples:**

```bash
curl http://localhost:3000/api/buildings
curl http://localhost:3000/api/buildings/defensive/cannon
```

### Troops — `/api/troops`

| Method | Path                              | Description                  |
|--------|-----------------------------------|------------------------------|
| GET    | `/api/troops`                     | List all troop categories    |
| GET    | `/api/troops/{category}`          | List troops in a category    |
| GET    | `/api/troops/{category}/{name}`   | Get a specific troop's data  |

**Categories:** `elixir`, `dark_elixir`, `super`

**Examples:**

```bash
curl http://localhost:3000/api/troops
curl http://localhost:3000/api/troops/elixir
```

## Configuration

All settings are controlled via environment variables. Copy `.env.example` to `.env` for reference.

| Variable        | Default       | Description                          |
|-----------------|---------------|--------------------------------------|
| `PORT`          | `3000`        | Server listen port                   |
| `ENVIRONMENT`   | `development` | `development` or `production`        |
| `LOG_LEVEL`     | `info`        | `debug`, `info`, `warn`, `error`     |
| `DATA_DIR`      | `data`        | Path to the JSON data directory      |
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
