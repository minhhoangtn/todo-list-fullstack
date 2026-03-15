# Todo List Backend

A REST API backend built with Go and Gin.

## Requirements

- Go 1.24+

## Project Structure

```text
backend/
├── cmd/
│   └── app/
│       └── main.go           # Application entry point & dependency wiring
├── internal/
│   ├── api/                  # API abstraction layer
│   │   ├── server.go         # Server interface (swap HTTP/GraphQL here)
│   │   └── http/             # Gin HTTP implementation
│   │       ├── server.go     # Route registration & server lifecycle
│   │       └── todo_handler.go
│   ├── services/             # Business logic
│   │   └── todo_service.go
│   ├── repositories/         # Data access layer
│   │   └── todo_repo.go
│   └── models/               # Data models
│       └── todo.go
├── go.mod
└── go.sum
```

### Swapping the API layer

To use a different HTTP framework or add GraphQL, implement the `api.Server` interface in a new sub-package and swap the one line in `main.go`:

```go
// Current: Gin HTTP
var server api.Server = httpApi.NewServer(svc)

// Future: GraphQL
var server api.Server = graphqlApi.NewServer(svc)
```

Business logic in `services/` and `repositories/` is never touched.

## Setup

Install dependencies:

```bash
go mod download
```

## Running

### Development (with hot reload)

```bash
go tool air
```

The server will restart automatically on code changes.

### Production

```bash
go run ./cmd/app
```

The server runs on `http://localhost:8080`.

## API

| Method | Endpoint | Description       |
| ------ | -------- | ----------------- |
| GET    | `/todos` | Get all todos     |
| POST   | `/todos` | Create a new todo |

### POST /todos

Request body:

```json
{
  "body": "Buy groceries",
  "completed": false
}
```
