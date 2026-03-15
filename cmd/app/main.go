package main

import (
	"minhhoangtn/todo-list-fullstack/internal/api"
	httpApi "minhhoangtn/todo-list-fullstack/internal/api/http"
	"minhhoangtn/todo-list-fullstack/internal/repositories"
	"minhhoangtn/todo-list-fullstack/internal/services"
)

func main() {
	repo := repositories.NewTodoRepository()
	svc := services.NewTodoService(repo)

	var server api.Server = httpApi.NewServer(svc)
	server.Start(":8080")
}
