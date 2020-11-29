# Seating Assignment App

This is a simple demo webapp.

- Users can create office locations
- Each location has one or more floorplans
- Each floorplan has several desks
- Users can be assigned to a desk
- Users can search for specific users and navigate locations

## Running

Requirements:

- `docker`, `docker-compose`. Can `brew cask install docker` on MacOS.
- Public Mapbox API token. Put it in `./.env` with `MAPBOX_TOKEN='...'`

Run dev servers: `./run start`

Go to http://localhost:8080

Changes to files in `backend/` will rebuild and restart the backend server

Changes to files in `frontend/src` will be hot-reloaded in the browser

## Info

The backend is a barebones golang HTTP server:

- https://github.com/rs/zerolog for structured logging
- https://github.com/gorilla/mux for HTTP routing
- https://github.com/alexflint/go-arg for parsing CLI arguments and environment variables
- https://github.com/jackc/pgx for communicating with postgres
- https://github.com/pkg/errors for better error handling
- Otherwise, uses Go standard library

Database is Postgres + PostGIS to enable efficient location-based lookup of office locations.

Migrations are automatically applied during application startup. This isn't my favorite approach for a number of reasons, but, it's fast and relatively easy at small scale.

The frontend is a pretty vanilla Webpack+React setup:

- Based on [my react-boilerplate project](https://github.com/austinhyde/react-boilerplate)
- Uses `semantic-ui-react` and `fomantic-ui-css` (semantic is KLO, fomantic is the fork/continuation) for base UI components
- Uses `react-map-gl` for a simple map

The frontend communicates with the backend via HTTP, using a plain old JSON-RPC style.

All builds and dev servers are dockerized, docker-compose is used to orchestrate dev servers and prod builds. This means _all_ dependencies are captured and hermetic, and the only hard dependency for a developer is docker.