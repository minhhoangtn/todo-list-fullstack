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
	s.router.GET("/todos", s.handler.GetTodos)
	s.router.POST("/todos", s.handler.CreateTodo)

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
