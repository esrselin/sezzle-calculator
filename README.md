# Sezzle Calculator

A full-stack calculator application built with React and TypeScript on the frontend and Go on the backend.

The frontend communicates with the backend through a REST API to perform arithmetic operations.

## Features

- Addition
- Subtraction
- Multiplication
- Division
- Exponentiation
- Square root
- Percentage
- Input validation
- Error handling
- Division by zero handling
- Invalid operation handling
- Backend connection error handling
- Responsive design
- Unit tests for frontend and backend
- Docker support

## Tech Stack

### Frontend

- React
- TypeScript
- Vite
- Vitest
- React Testing Library

### Backend

- Go
- Go standard library
- REST API

### Development Tools

- Docker
- Docker Compose

## Project Structure

```text
sezzle-calculator/
├── backend/
│   ├── main.go
│   ├── main_test.go
│   ├── go.mod
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── App.tsx
│   │   ├── App.css
│   │   ├── App.test.tsx
│   │   └── test/
│   │       └── setup.ts
│   ├── package.json
│   ├── package-lock.json
│   ├── vite.config.ts
│   └── Dockerfile
├── .dockerignore
├── .gitignore
├── docker-compose.yml
├── AI_USAGE.md
└── README.md
```

## Getting Started

### Prerequisites

For local development:

- Go
- Node.js
- npm

For Docker:

- Docker Desktop
- Docker Compose

## Running Locally

### Backend

```bash
cd backend
go run main.go
```

The backend starts on `http://localhost:8080`.

API endpoint:

```text
POST /calculate
```

### Frontend

Open another terminal:

```bash
cd frontend
npm install
npm run dev
```

The frontend is available at `http://localhost:5173`.

## Running with Docker

From the project root:

```bash
docker compose up --build
```

This starts both services:

- Frontend: `http://localhost:5173`
- Backend: `http://localhost:8080`

Stop the containers with:

```bash
docker compose down
```

## API Usage

### POST `/calculate`

The API accepts two numbers and an operation.

Example:

```json
{
  "firstNumber": 10,
  "operation": "+",
  "secondNumber": 5
}
```

Response:

```json
{
  "result": 15
}
```

### Supported Operations

| Operation | Description    | Example       |
| --------- | -------------- | ------------- |
| `+`       | Addition       | `10 + 5 = 15` |
| `-`       | Subtraction    | `10 - 5 = 5`  |
| `*`       | Multiplication | `10 * 5 = 50` |
| `/`       | Division       | `10 / 5 = 2`  |
| `^`       | Exponentiation | `2 ^ 3 = 8`   |
| `sqrt`    | Square root    | `√25 = 5`     |
| `%`       | Percentage     | `50 % 10 = 5` |

Percentage is calculated as:

```text
firstNumber × secondNumber / 100
```

For square root, only the first number is used.

For exponentiation, the first number is raised to the power of the second number.

## Error Responses

### Division by zero

```json
{
  "error": "Cannot divide by zero."
}
```

### Invalid operation

```json
{
  "error": "invalid operation"
}
```

### Invalid request body

```json
{
  "error": "Invalid request body"
}
```

### Square root of a negative number

```json
{
  "error": "Cannot calculate square root of a negative number."
}
```

## Validation and Error Handling

The application handles:

- Missing numbers
- Invalid numeric input
- Invalid JSON request bodies
- Division by zero
- Square root of negative numbers
- Unsupported operations
- Unsupported HTTP methods
- Backend connection errors

The backend returns JSON error responses with appropriate HTTP status codes.

The frontend validates required input before making an API request and displays backend or connection errors to the user.

## Testing

### Backend

Run tests:

```bash
cd backend
go test -v
```

Run coverage:

```bash
go test -cover
```

Current backend coverage is approximately 77%.

Backend tests cover:

- Basic arithmetic operations
- Exponentiation
- Square root
- Percentage
- Negative numbers
- Decimal numbers
- Division by zero
- Invalid operations
- Invalid JSON
- Unsupported HTTP methods
- CORS preflight requests

### Frontend

Run tests:

```bash
cd frontend
npm test -- --run
```

Run coverage:

```bash
npm test -- --run --coverage
```

Frontend tests cover:

- Calculator rendering
- Successful calculations
- Missing input validation
- API error handling
- Square root
- Percentage

The current frontend test suite contains 6 tests.

## Production Build

```bash
cd frontend
npm run build
```

The project builds successfully with TypeScript and Vite.

## Design Decisions

### Separate Frontend and Backend

The frontend handles user interaction and input validation, while the backend handles calculation logic.

This keeps the calculation logic independent from the UI and allows the API to be used separately.

### Centralized Calculation Logic

The backend uses a dedicated `calculate` function for arithmetic operations.

The HTTP handler is responsible for:

1. Validating the HTTP method
2. Decoding the request
3. Calling the calculation function
4. Returning a JSON response

This keeps HTTP handling separate from calculation logic and makes the core functionality easier to test.

### Error Handling

Calculation errors are returned from the calculation layer and converted into JSON API responses by the HTTP handler.

The frontend also handles validation errors and backend connection errors.

### Responsive UI

Responsive CSS keeps the calculator usable on smaller screens while keeping the interface simple.

### Docker

The frontend and backend are containerized separately and orchestrated with Docker Compose.

This keeps the services isolated while providing a simple way to run the complete application with one command.

## Assumptions

- Percentage is calculated as `firstNumber × secondNumber / 100`.
- Square root uses only the first number.
- The second number is not required for square root.
- Floating-point numbers are used to support decimal calculations.
- The backend runs on port `8080`.
- The frontend development server runs on port `5173`.
- The frontend connects to the backend using `http://localhost:8080`.

## AI Usage

AI tools were used as a development assistant for debugging, test design, code review, and documentation.

The prompts used during development are documented separately in [AI_USAGE.md](AI_USAGE.md).
