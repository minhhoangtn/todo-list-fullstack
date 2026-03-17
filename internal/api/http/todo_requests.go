package http

type CreateTodoRequest struct {
	Body      string `json:"body"      binding:"required"`
	Completed bool   `json:"completed"`
}

type UpdateTodoRequest struct {
	Body      string `json:"body"      binding:"required"`
	Completed bool   `json:"completed"`
}
