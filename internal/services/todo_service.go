package services

import (
	"errors"
	"minhhoangtn/todo-list-fullstack/internal/errs"
	"minhhoangtn/todo-list-fullstack/internal/models"
	"minhhoangtn/todo-list-fullstack/internal/repositories"
)

type TodoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAll() []models.Todo {
	return s.repo.GetAll()
}

func (s *TodoService) Create(todo models.Todo) models.Todo {
	return s.repo.Add(todo)
}

func (s *TodoService) Update(id string, todo models.Todo) (models.Todo, error) {
	result, err := s.repo.Update(id, todo)
	if errors.Is(err, models.ErrTodoNotFound) {
		return models.Todo{}, errs.NotFound(err)
	}

	if err != nil {
		return models.Todo{}, errs.Internal(err)
	}
	return result, nil
}

func (s *TodoService) Delete(id string) error {
	err := s.repo.Delete(id)
	if errors.Is(err, models.ErrTodoNotFound) {
		return errs.NotFound(err)
	}

	if err != nil {
		return errs.Internal(err)
	}

	return nil
}
