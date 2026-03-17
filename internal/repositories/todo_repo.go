package repositories

import (
	"minhhoangtn/todo-list-fullstack/internal/models"
)

type TodoRepository interface {
	GetAll() []models.Todo
	Add(todo models.Todo) models.Todo
	Update(id string, todo models.Todo) (models.Todo, error)
	Delete(id string) error
}
