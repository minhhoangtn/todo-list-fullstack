// @title           Todo List API
// @version         1.0
// @description     A simple todo list REST API.

// @host      localhost:8080
// @BasePath  /

package main

import (
	"log"
	"minhhoangtn/todo-list-fullstack/internal/api"
	httpApi "minhhoangtn/todo-list-fullstack/internal/api/http"
	"minhhoangtn/todo-list-fullstack/internal/repositories"
	"minhhoangtn/todo-list-fullstack/internal/services"

	_ "minhhoangtn/todo-list-fullstack/docs"
)

func main() {
	repo := repositories.NewTodoLocalRepository()
	svc := services.NewTodoService(repo)

	var server api.Server = httpApi.NewServer(svc)
	log.Fatal(server.Start(":8080"))
}
