package repositories

import (
	"minhhoangtn/todo-list-fullstack/internal/models"

	"github.com/google/uuid"
)

type TodoLocalRepository struct {
	todos []models.Todo
}

func NewTodoLocalRepository() TodoRepository {
	return &TodoLocalRepository{todos: []models.Todo{}}
}

func (r *TodoLocalRepository) GetAll() []models.Todo {
	return r.todos
}

func (r *TodoLocalRepository) Add(todo models.Todo) models.Todo {
	todo.ID = uuid.New().String()
	r.todos = append(r.todos, todo)
	return todo
}

func (r *TodoLocalRepository) Update(id string, todo models.Todo) (models.Todo, error) {
	for i, t := range r.todos {
		if t.ID == id {
			r.todos[i].Body = todo.Body
			r.todos[i].Completed = todo.Completed
			return r.todos[i], nil
		}
	}
	return models.Todo{}, models.ErrTodoNotFound
}

func (r *TodoLocalRepository) Delete(id string) error {
	for i, t := range r.todos {
		if t.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return models.ErrTodoNotFound
}
