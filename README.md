# Seating Assignment App

This is a simple demo webapp.

- Users can create office locations and view them on a global map
- Users can create one or more floorplans for each location, and pan/zoom around them
- Each floorplan has several desks, and users can be assigned to those desks
- Users can search for specific people to see where they sit

This project serves multiple purposes for me:

- It's a good small-scale representation of the non-public web application work I've done, in terms of application architecture and various machinery around it
- It's a low-stakes application with a well-defined feature set that I can experiment with different architectures/frameworks/languages/etc in. I have a number of different "experiments" I want to mess with or refine over time
- It's a good baseline/boilerplate for other applications. Once I settle on specific implementation details, I hope to extract the core boilerplate into a standalone repo for copy+paste in new applications

To do list:
- [x] Set up docker development environment
- [x] Set up migrations, core db schema
- [x] Set up rough backend framework, a few example API endpoints
- [x] Set up rough frontend structure, some basic interaction
- [ ] Implement all API endpoints
- [ ] Implement all core functionality in UI
- [ ] Refine API -> Service -> DB layers and responsibilities
- [ ] Refine UI state management
- [ ] Refine database change management approach
- [ ] ...?

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