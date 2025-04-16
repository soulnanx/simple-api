package server

import (
	"net/http"

	"simple-api/internal/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	s := &Server{router: http.NewServeMux()}
	s.routes()
	return s
}

func (s *Server) routes() {
	// Swagger
	s.router.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Handlers
	taskHandler := handlers.NewTaskHandler()
	s.router.HandleFunc("/tasks", taskHandler.HandleTasks)
	s.router.HandleFunc("/ping", taskHandler.Ping)
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
