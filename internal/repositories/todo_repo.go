package repositories

import "minhhoangtn/todo-list-fullstack/internal/models"

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
	todo.ID = len(r.todos) + 1
	r.todos = append(r.todos, todo)
	return todo
}
