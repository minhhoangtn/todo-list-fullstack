package repositories

import (
	"errors"
	"minhhoangtn/todo-list-fullstack/internal/models"

	"github.com/google/uuid"
)

type TodoRepository struct {
	todos []models.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) GetAll() []models.Todo {
	return r.todos
}

func (r *TodoRepository) Add(todo models.Todo) models.Todo {
	todo.ID = uuid.New().String()
	r.todos = append(r.todos, todo)
	return todo
}

func (r *TodoRepository) Update(id string, todo models.Todo) (models.Todo, error) {
	for i, t := range r.todos {
		if t.ID == id {
			r.todos[i].Body = todo.Body
			r.todos[i].Completed = todo.Completed
			return r.todos[i], nil
		}
	}
	return models.Todo{}, errors.New("todo not found")
}

func (r *TodoRepository) Delete(id string) error {
	for i, t := range r.todos {
		if t.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
