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

func (s *TodoService) Update(id string, todo models.Todo) (models.Todo, error) {
	return s.repo.Update(id, todo)
}

func (s *TodoService) Delete(id string) error {
	return s.repo.Delete(id)
}
