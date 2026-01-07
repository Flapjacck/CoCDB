# CoCDB

clash of clans api database with a simple REST API.

## Setup

### Prerequisites

- Go 1.25.5 or higher
- Git

### Installation

1. Clone the repository:

```bash
git clone https://github.com/flapjacck/CoCDB.git
cd cocdb
```

1. Install dependencies:

```bash
go mod download
```

## Running the API

Start the server:

```bash
go run main.go
```

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /api/data/{path}` - Retrieve JSON data from the `data/` directory
  - Example: `GET /api/data/home_village/buildings/defensive/cannon`
  - Example: `GET /api/data/home_village/troops/elixir/barbarian`

## Attribution & Licensing

We use material from the Clash of Clans Wiki at Fandom to make this database. Please see their licensing terms below:

Community content on Fandom is available under the Creative Commons Attribution-ShareAlike license (CC BY‑SA) unless otherwise noted. See: <https://clashofclans.fandom.com/wiki/Clash_of_Clans_Wiki>

Some images or other media may have separate licenses; always check the media page on Fandom before reuse.

All JSON files include a `source`, `source_url`, and `source_license` field pointing to the original wiki article and license.
