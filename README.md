# Calculator

A React calculator backed by a Go expression parser and evaluator.

## Run with Docker

Install Docker with Docker Compose, then run from the project root:

```sh
docker compose up --build -d
```

Open http://localhost:3000. Nginx serves the frontend and forwards `/calculate`
to the backend through the Compose network. The backend does not require a
published host port. Both services include health checks; the frontend starts
after the backend is healthy.

To use a different frontend port:

```sh
FRONTEND_PORT=8000 docker compose up --build -d
```

Inspect status and logs:

```sh
docker compose ps
docker compose logs -f
```

Stop the services:

```sh
docker compose down
```

## Local development

Start the backend in one terminal:

```sh
cd backend
go run .
```

Start the frontend in another terminal:

```sh
cd frontend
corepack pnpm install --frozen-lockfile
corepack pnpm dev
```

Open the Vite URL printed in the terminal. Keep `VITE_API_URL` empty to use the
Vite proxy to the backend on port 8080.
