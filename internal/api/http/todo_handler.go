package http

import (
	"minhhoangtn/todo-list-fullstack/internal/models"
	"minhhoangtn/todo-list-fullstack/internal/services"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	c.JSON(200, h.service.GetAll())
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, h.service.Create(todo))
}
