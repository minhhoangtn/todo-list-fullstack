# Todo List Backend

A REST API backend built with Go and Gin.

## Requirements

- Go 1.24+

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
go run main.go
```

The server runs on `http://localhost:8080`.

## API

| Method | Endpoint | Description  |
| ------ | -------- | ------------ |
| GET    | `/ping`  | Health check |
