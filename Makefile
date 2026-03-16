APP     := ./cmd/app
DOCS    := docs

.PHONY: run dev build swag clean tidy

run:
	go run $(APP)

dev:
	go tool air

swag:
	go tool swag init -g $(APP)/main.go -o $(DOCS)

