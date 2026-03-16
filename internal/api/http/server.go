package http

import (
	"minhhoangtn/todo-list-fullstack/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	todos := s.router.Group("/todos")
	todos.GET("", s.handler.GetTodos)
	todos.POST("", s.handler.CreateTodo)
	todos.PATCH("/:id", s.handler.UpdateTodo)
	todos.DELETE("/:id", s.handler.DeleteTodo)

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
