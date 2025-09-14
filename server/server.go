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

	// middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)    // логирование запросов
	r.Use(middleware.Recoverer) // защита от паник

	// маршруты
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", s.handlers.HandleCreateTask)

		r.Get("/", s.handlers.HandleGetAllTasks) // все задачи или фильтр по completed
		r.Get("/{id}", s.handlers.HandleGetTask) // задача по ID

		r.Patch("/", s.handlers.HandleUpdateTask)  // ?id={id}&completed={completed}
		r.Delete("/", s.handlers.HandleDeleteTask) // ?id={id}
	})

	log.Println("Зарегистрированные маршруты:")
	log.Println("POST   /tasks                  - создание задачи")
	log.Println("GET    /tasks                  - получение всех задач")
	log.Println("GET    /tasks?completed=false  - получение незавершенных задач")
	log.Println("GET    /tasks/{id}             - получение задачи по ID")
	log.Println("PATCH  /tasks?id={id}&completed={completed} - обновление задачи")
	log.Println("DELETE /tasks?id={id}          - удаление задачи")
	log.Printf("Сервер запускается на порту %s", port)

	return http.ListenAndServe(":"+port, r)
}
