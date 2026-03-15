package http

import (
	"minhhoangtn/todo-list-fullstack/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	handler *TodoHandler
}

func NewServer(svc *services.TodoService) *Server {
	s := &Server{
		router:  gin.Default(),
		handler: NewTodoHandler(svc),
	}
	s.registerRoutes()
	return s
}

func (s *Server) registerRoutes() {
	s.router.GET("/todos", s.handler.GetTodos)
	s.router.POST("/todos", s.handler.CreateTodo)
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
