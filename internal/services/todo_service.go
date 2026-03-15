package services

import (
	"minhhoangtn/todo-list-fullstack/internal/models"
	"minhhoangtn/todo-list-fullstack/internal/repositories"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAll() []models.Todo {
	return s.repo.GetAll()
}

func (s *TodoService) Create(todo models.Todo) models.Todo {
	return s.repo.Add(todo)
}
