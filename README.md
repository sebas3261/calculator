# Calculator

A full-stack calculator built with **React, TypeScript, and Go**, featuring a custom mathematical expression parser and evaluator.

The application supports standard arithmetic operations, mathematical precedence, parentheses, exponentiation, square roots, and percentages.

## Tech Stack

| Layer | Technologies |
|---|---|
| Frontend | React, TypeScript, Vite |
| Backend | Go |
| Testing | Go testing, Vitest, React Testing Library |
| Deployment | Docker, Docker Compose, Nginx |

## Features

- Addition, subtraction, multiplication, and division.
- Exponentiation (`^`) and square roots (`√`).
- Postfix percentages (`%`).
- Parenthesized expressions and operator precedence.
- Keyboard input and interactive calculator controls.
- Expression editing, deletion, and clearing.
- Continuation from a previous result.
- Responsive layout for desktop and mobile.
- Error handling for invalid expressions and mathematical operations.

## Getting Started

### Option 1 — Docker (recommended)

**Requirements:** Docker with Docker Compose.

From the project root, run:

```bash
docker compose up --build -d
```

Open **http://localhost:3000**.

Nginx serves the frontend and forwards requests to `/calculate` through the Docker Compose network. The backend does not require a published host port.

Both services include health checks, and the frontend starts after the backend is healthy.

To use a different frontend port:

```bash
FRONTEND_PORT=8000 docker compose up --build -d
```

Check the container status:

```bash
docker compose ps
```

View logs:

```bash
docker compose logs -f
```

Stop the application:

```bash
docker compose down
```

### Option 2 — Local development

**Requirements:** Go 1.27.3 or newer, Node.js 24 or newer, and pnpm.

The Docker build uses Node.js 24 and pnpm 10.11.0. If Corepack is not available locally, use an installed pnpm CLI in place of `corepack pnpm`.

Start the backend in one terminal:

```bash
cd backend
go run .
```

The API runs on `http://localhost:8080`.

Start the frontend in another terminal:

```bash
cd frontend
corepack pnpm install --frozen-lockfile
corepack pnpm dev
```

Open the Vite URL printed in the terminal.

Optionally copy `.env.example` to `.env` for local configuration. Local `.env` files are ignored by Git. Leave `VITE_API_URL` empty to use the Vite development proxy, which forwards API requests to the Go backend on port 8080.

## Architecture

### Frontend

The frontend follows a lightweight **View/ViewModel** pattern.

```text
frontend/src/
├── lib/
│   └── http.ts
├── models/
├── services/
├── viewmodels/
└── views/
```

Responsibilities:

- **View:** Renders the calculator and handles presentation.
- **ViewModel:** Manages expression state, user interactions, loading states, and errors.
- **Services:** Provides calculator-specific API operations.
- **HTTP client:** Encapsulates reusable HTTP requests and error handling.
- **Models:** Defines TypeScript data contracts.

This separation keeps UI rendering independent of API communication and calculator state management.

### Backend

The Go backend separates expression parsing from mathematical evaluation.

```text
HTTP Handler
     |
Calculator Service
     |
     +-- Parser
     |     |
     |     +-- AST
     |
     +-- Evaluator
```

The parser converts a mathematical expression into an **abstract syntax tree (AST)**.

The evaluator traverses the AST and computes the result.

The calculator service coordinates both components, while the HTTP layer handles JSON requests and responses.

Parser and evaluator implementations are connected through interfaces, keeping their responsibilities separate and testable.

## API

### `POST /calculate`

Evaluates a mathematical expression.

**Request**

```json
{
  "expression": "2 + 3 * 4"
}
```

**Successful response — HTTP 200**

```json
{
  "answer": 14
}
```

**Error response — HTTP 422**

```json
{
  "error": "division by zero"
}
```

Example using `curl` while running the backend locally:

```bash
curl -X POST http://localhost:8080/calculate \
  -H "Content-Type: application/json" \
  -d '{"expression":"2 + 3 * 4"}'
```

When running through Docker, use the frontend's Nginx endpoint instead:

```bash
curl -X POST http://localhost:3000/calculate \
  -H "Content-Type: application/json" \
  -d '{"expression":"√9 + 2^3"}'
```

## Mathematical Behavior

The calculator respects operator precedence and parentheses.

| Expression | Result |
|---|---:|
| `2 + 3 * 4` | 14 |
| `(2 + 3) * 4` | 20 |
| `2^3^2` | 512 |
| `-2^2` | -4 |
| `2^-3` | 0.125 |
| `1e+21 * 2` | 2e+21 |
| `√(9 + 16)` | 5 |
| `50% * 200` | 100 |
| `25% + 12` | 12.25 |

### Percentage semantics

The `%` operator divides its operand by 100.

For example, `200 + 10%` evaluates to `200.1`, not `220`.

This differs from calculators that apply percentages contextually.

### Invalid expressions

Malformed expressions and invalid mathematical operations, such as division by zero or square roots of negative numbers, produce error responses rather than numerical results.

## Testing

### Backend

Run the Go tests:

```bash
cd backend
go test ./...
```

Run tests with the race detector and coverage reporting:

```bash
go test ./... -race -coverprofile=coverage.out
go tool cover -func=coverage.out
```

During development, the backend test suite reported **96.8% overall statement coverage**.

### Frontend

From the `frontend` directory:

```bash
corepack pnpm test
```

For coverage:

```bash
corepack pnpm test:coverage
```

During development, the frontend test suite reported **32 passing tests** and **94.73% statement coverage** over the configured files.

These results are historical measurements and may change as the implementation evolves.

## AI Assistance

AI tools were used during development for technical discussion, code implementation, testing, review, and documentation.

The project was developed using a combination of manually established architecture, human-directed implementation decisions, and AI-assisted development.

For details about the tools used, their contributions, and representative prompts, see [AI Assistance Disclosure](./AI_USAGE.md).

## Author

Sebastián Sánchez Sandoval

[Portfolio](https://sebas3261.com) · [GitHub](https://github.com/sebas3261)