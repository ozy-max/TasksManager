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
	r := chi.NewRouter()

	r.Use(middleware.RequestID) // уникальный ID для каждого запроса
	r.Use(middleware.RealIP)    // определение реального IP клиента
	r.Use(middleware.Logger)    // логирование запросов
	r.Use(middleware.Recoverer) // защита от паник

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", s.handlers.HandleCreateTask) // создание задачи

		r.Get("/", s.handlers.HandleGetAllTasks) // все задачи или фильтр по completed
		r.Get("/{id}", s.handlers.HandleGetTask) // задача по ID

		r.Patch("/", s.handlers.HandleUpdateTask)  // ?id={id}&completed={completed}
		r.Delete("/", s.handlers.HandleDeleteTask) // ?id={id}
	})

	log.Printf("Сервер запускается на порту %s", port)
	return http.ListenAndServe(":"+port, r)
}
