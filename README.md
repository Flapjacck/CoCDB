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
- `GET /api/data/{path}` - Retrieve JSON data from the data directory
  - Example: `GET /api/data/home_village/buildings/defensive/cannon`
  - Example: `GET /api/data/home_village/troops/elixir/barbarian`
