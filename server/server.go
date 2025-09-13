package server

import (
	"log"
	"net/http"
	handlers "tasks_manager/httpUtils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	handlers *handlers.HTTPHandlers
}

func NewServer(handlers *handlers.HTTPHandlers) *Server {
	return &Server{handlers: handlers}
}

func (s *Server) StartServer(port string) error {
	router := chi.NewRouter()

	// Добавляем middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	// Настройка маршрутов
	router.Route("/tasks", func(r chi.Router) {
		r.Post("/", s.handlers.HandleCreateTask)            // POST /tasks
		r.Get("/", s.handlers.HandleGetAllTasks)            // GET /tasks
		r.Get("/", s.handlers.HandleGetAllUncompletedTasks) // GET /tasks?completed=false
		r.Get("/{id}", s.handlers.HandleGetTask)            // GET /tasks/{id}
		r.Patch("/{id}", s.handlers.HandleCompleteTask)     // PATCH /tasks/{id}
		r.Delete("/{id}", s.handlers.HandleDeleteTask)      // DELETE /tasks/{id}
	})

	log.Printf("Сервер запускается на порту %s", port)
	return http.ListenAndServe(":"+port, router)
}
