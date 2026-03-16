package http

import (
	"minhhoangtn/todo-list-fullstack/internal/models"
	"minhhoangtn/todo-list-fullstack/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

// GetTodos godoc
// @Summary      List all todos
// @Tags         todos
// @Produce      json
// @Success      200  {array}   models.Todo
// @Router       /todos [get]
func (h *TodoHandler) GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetAll())
}

// CreateTodo godoc
// @Summary      Create a todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      models.Todo  true  "Todo body"
// @Success      201   {object}  models.Todo
// @Failure      400   {object}  map[string]string
// @Router       /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, h.service.Create(todo))
}
