# Golang-Backend Starter

---

## Tech Stack

- Go
- Echo for HTTP routing
- GORM for database access
- PostgreSQL as the database
- JWT for authentication
- Validator for request validation
- godotenv for loading `.env` files

## Project Structure

```text
go-backend-boilerplate/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── auth/                   # JWT token create/validate logic
│   ├── config/                 # Environment and database config
│   ├── domain/
│   │   └── user/               # User auth/profile feature
│   │       └── dto/            # User request/response DTOs
│   ├── httpresponse/           # Common error response shape
│   ├── middlewares/            # Auth middleware
│   └── server/                 # Echo server setup
├── .air.toml                   # Air config for live reload
├── .env.example                # Example environment variables
├── go.mod                      # Go module and dependencies
└── go.sum
```

## How The Code Flows

For most features, the code follows this pattern:

```text
Route -> Handler -> Service -> Repository -> Database
```

- **Route** decides the URL and HTTP method.
- **Handler** reads the request and returns the response.
- **Service** contains the main business logic.
- **Repository** talks to the database.
- **DTO** defines request and response data shapes.

## Requirements

Before running the project, install:

- Go
- PostgreSQL
- Git

You can check your Go version with:

```bash
go version
```

The `DSN` value tells GORM how to connect to PostgreSQL.

## Run The Project

Install dependencies:

```bash
go mod tidy
```

Start the server:

```bash
go run cmd/main.go
```

Or, if you are using Air for live reload:

```bash
air
```

If everything is okay, the server will start on:

```text
http://localhost:8080
```

Check the health route:

```bash
curl http://localhost:8080/health
```

Expected response:

```text
running
```

## Build The Project

Create a local binary:

```bash
go build -o bin/go-backend-boilerplate ./cmd/main.go
```

Run the binary:

```bash
./bin/go-backend-boilerplate
```

For a smaller production-style binary:

```bash
CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/go-backend-boilerplate ./cmd/main.go
```

Before building for production, it is a good habit to run:

```bash
go mod tidy
go test ./...
go vet ./...
```

## API Routes

### Auth Routes

Register a new user:

```http
POST /api/v1/auth/register
```

Example:

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "mamun",
    "email": "mamun@example.com",
    "password": "secret123"
  }'
```

Login:

```http
POST /api/v1/auth/login
```

Example:

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "mamun@example.com",
    "password": "secret123"
  }'
```

The login response returns a JWT token. Keep that token for protected routes.

Get current user:

```http
GET /api/v1/auth/me
```

Example:

```bash
curl http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Important Concepts

### `cmd/main.go`

This is where the application starts. It loads environment variables, connects to the database, and starts the server.

### `internal/config`

This package reads values from `.env` and creates the database connection.

### `internal/server`

This package creates the Echo app, adds middleware, registers routes, and starts listening on the selected port.

### `internal/domain`

This folder contains the main business domains: user, event, and booking. Each domain has its own handler, service, repository, entity, route registration, and DTOs.

### Handler

A handler receives an HTTP request. It usually does these things:

1. Reads JSON input
2. Validates the input
3. Calls the service
4. Sends JSON response

### Service

A service contains business logic. For example, the booking service checks if enough tickets are available before creating a booking.

### Repository

A repository is responsible for database queries. This keeps database code separate from business logic.

### Middleware

Middleware runs before the handler. In this project, auth middleware checks the JWT token and adds user information to the request context.

## Database Tables

The project uses GORM `AutoMigrate`, so tables are created automatically when the server starts.

Current main tables:

- users

## Common Problems

### `.env` file not found

Make sure you created a `.env` file:

```bash
cp .env.example .env
```

### Database connection failed

Check these things:

- PostgreSQL is running
- Database name is correct
- Username and password are correct
- Port is correct, usually `5432`

### Protected route says unauthorized

Make sure the request has this header:

```text
Authorization: Bearer YOUR_TOKEN_HERE
```

## Also make sure the token comes from the login route.

---

- `go mod init <module-path>` (progmamun/go-backend-boilerplate)
- `go mod tidy`
- `go get github.com/joho/godotenv`
- `go get github.com/labstack/echo/v5`
- `air init`
- `go build -o bin/foldername ./cmd`
