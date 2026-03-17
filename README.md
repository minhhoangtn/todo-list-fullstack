# Todo List Backend

A REST API backend built with Go and Gin.

## Requirements

- Go 1.24+

## Project Structure

```text
backend/
├── cmd/
│   └── app/
│       └── main.go              # Application entry point & dependency wiring
├── docs/                        # Auto-generated Swagger docs (do not edit manually)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── api/                     # API abstraction layer
│   │   ├── server.go            # Server interface (swap HTTP/GraphQL here)
│   │   └── http/                # Gin HTTP implementation
│   │       ├── server.go        # Route registration & server lifecycle
│   │       ├── todo_handler.go  # Request handlers with Swagger annotations
│   │       ├── todo_requests.go # Request body structs
│   │       └── error_handler.go # AppError → HTTP status mapping & validation errors
│   ├── services/                # Business logic
│   │   └── todo_service.go
│   ├── repositories/            # Data access layer
│   │   ├── todo_repo.go         # TodoRepository interface
│   │   └── todo_local_repo.go   # In-memory implementation (UUID IDs)
│   ├── models/                  # Shared data types
│   │   ├── todo.go
│   │   └── errors.go            # Sentinel errors (e.g. ErrTodoNotFound)
│   └── errs/                    # Application error types
│       └── app_error.go         # AppError type, error codes & constructors
├── go.mod
└── go.sum
```

### Error handling

Errors flow through layers with clear separation:

- **Repository**: returns sentinel errors from `models/errors.go` (e.g. `models.ErrTodoNotFound`)
- **Service**: wraps them into `*errs.AppError` with a machine-readable code and safe user message
- **HTTP handler**: `handleAppError` in `error_handler.go` maps `AppError.Code` to the correct HTTP status

### Swapping the API layer

To use a different HTTP framework or add GraphQL, implement the `api.Server` interface in a new sub-package and swap one line in `main.go`:

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

## Makefile commands

The server runs on `http://localhost:8080`.
Interactive API documentation is available at:

```text
http://localhost:8080/swagger/index.html
```

| Command     | Description                                                         |
| ----------- | ------------------------------------------------------------------- |
| `make run`  | Run the app, The server will restart automatically on code changes. |
| `make dev`  | Run with hot reload (air)                                           |
| `make swag` | Regenerate Swagger docs                                             |
