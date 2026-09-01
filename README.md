````md
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
- Responsive design
- Unit tests for frontend and backend

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

## Project Structure

```text
sezzle-calculator/
├── backend/
│   ├── main.go
│   ├── main_test.go
│   └── go.mod
│
├── frontend/
│   ├── src/
│   │   ├── App.tsx
│   │   ├── App.css
│   │   ├── App.test.tsx
│   │   └── test/
│   │       └── setup.ts
│   ├── package.json
│   └── vite.config.ts
│
├── .gitignore
├── AI_USAGE.md
└── README.md
```
````

## Getting Started

### Prerequisites

Make sure the following are installed:

- Go
- Node.js
- npm

### Clone the Repository

```bash
git clone https://github.com/esrselin/sezzle-calculator.git
cd sezzle-calculator
```

## Running the Backend

Open a terminal in the `backend` directory:

```bash
cd backend
go run main.go
```

The backend starts on:

```text
http://localhost:8080
```

The calculator API endpoint is:

```text
POST /calculate
```

## Running the Frontend

Open another terminal in the `frontend` directory:

```bash
cd frontend
npm install
npm run dev
```

The frontend is available at:

```text
http://localhost:5173
```

The frontend communicates with the backend running on port `8080`.

## API Usage

### POST `/calculate`

The API accepts two numbers and an operation.

### Addition

Request:

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

For percentage, the calculation is:

```text
firstNumber × secondNumber / 100
```

For example:

```text
50 % 10 = 50 × 10 / 100 = 5
```

For square root, only the first number is used.

### Error Responses

Division by zero:

```json
{
  "error": "Cannot divide by zero."
}
```

Invalid operation:

```json
{
  "error": "invalid operation"
}
```

Invalid request body:

```json
{
  "error": "Invalid request body"
}
```

Square root of a negative number:

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

## Testing

### Backend

Run the backend tests:

```bash
cd backend
go test -v
```

Run backend coverage:

```bash
go test -cover
```

The current backend coverage is approximately 77%.

The backend tests cover:

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

Run the frontend tests:

```bash
cd frontend
npm test -- --run
```

Run frontend tests with coverage:

```bash
npm test -- --run --coverage
```

The frontend tests cover:

- Calculator rendering
- Successful calculations
- Missing input validation
- API error handling
- Square root
- Percentage
- Exponentiation

## Production Build

To create a production build:

```bash
cd frontend
npm run build
```

The project builds successfully with TypeScript and Vite.

## Design Decisions

### Separate Frontend and Backend

The frontend handles user interaction and input validation, while the backend handles the calculation logic.

This keeps the calculation logic independent from the UI and allows the API to be used separately from the frontend.

### Centralized Calculation Logic

The backend uses a dedicated `calculate` function for arithmetic operations.

The HTTP handler is responsible for:

1. Validating the HTTP method
2. Decoding the request
3. Calling the calculation function
4. Returning a JSON response

This keeps HTTP handling separate from the calculation logic and makes the core functionality easier to test.

### Error Handling

Calculation errors are returned from the calculation layer and converted into JSON API responses by the HTTP handler.

The frontend also handles validation errors and backend connection errors so that failures are displayed to the user.

### Responsive UI

The calculator uses responsive CSS to remain usable on smaller screens while keeping the interface simple.

## Assumptions

- Percentage is calculated as `firstNumber × secondNumber / 100`.
- Square root uses only the first number.
- The second number is not required for square root.
- Floating-point numbers are used to support decimal calculations.
- The backend runs on port `8080`.
- The frontend development server runs on port `5173`.
- The frontend currently connects to the backend using `http://localhost:8080`.

## AI Usage

AI tools were used as a development assistant for debugging, test design, code review, and documentation.

The prompts used during development are documented separately in [AI_USAGE.md](AI_USAGE.md).

```

```
