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
// @Param        todo  body      CreateTodoRequest  true  "Todo to create"
// @Success      201   {object}  models.Todo
// @Failure      400   {object}  map[string]string
// @Router       /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, handleValidationErrors(err))
		return
	}
	todo := models.Todo{Body: req.Body, Completed: req.Completed}
	c.JSON(http.StatusCreated, h.service.Create(todo))
}

// UpdateTodo godoc
// @Summary      Update a todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path      string             true  "Todo ID"
// @Param        todo  body      UpdateTodoRequest  true  "Updated todo fields"
// @Success      200   {object}  models.Todo
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /todos/{id} [patch]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var req UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, handleValidationErrors(err))
		return
	}
	updated, err := h.service.Update(id, models.Todo{Body: req.Body, Completed: req.Completed})
	if err != nil {
		handleAppError(c, err)
		return
	}
	c.JSON(http.StatusOK, updated)
}

// DeleteTodo godoc
// @Summary      Delete a todo
// @Tags         todos
// @Param        id   path  string  true  "Todo ID"
// @Success      204
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		handleAppError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
