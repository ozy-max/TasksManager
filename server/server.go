package server

import (
	"log"
	"net/http"
	handlers "tasks_manager/httpUtils"

	"github.com/gorilla/mux"
)

type Server struct {
	handlers *handlers.HTTPHandlers
}

func NewServer(handlers *handlers.HTTPHandlers) *Server {
	return &Server{handlers: handlers}
}

func (s *Server) StartServer(port string) error {
	router := mux.NewRouter()

	// Регистрируем маршруты
	router.Path("/tasks").Methods("POST").HandlerFunc(s.handlers.HandleCreateTask)

	router.Path("/tasks").Methods("GET").HandlerFunc(s.handlers.HandleGetAllTasks)

	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.handlers.HandleGetAllUncompletedTasks)

	router.Path("/tasks/{id}").Methods("GET").HandlerFunc(s.handlers.HandleGetTask)

	router.Path("/tasks").Methods("PATCH").
		Queries("id", "{id}", "completed", "{completed}").
		HandlerFunc(s.handlers.HandleUpdateTask)

	router.Path("/tasks").Methods("DELETE").Queries("id", "{id}").HandlerFunc(s.handlers.HandleDeleteTask)

	// Debug: добавляем логирование всех маршрутов
	log.Println("Зарегистрированные маршруты:")
	log.Println("POST /tasks - создание задачи")
	log.Println("GET /tasks - получение всех задач")
	log.Println("GET /tasks?completed=false - получение незавершенных задач")
	log.Println("GET /tasks/{id} - получение задачи по ID")
	log.Println("PATCH /tasks/{id}/{completed} - обновление задачи")
	log.Println("DELETE /tasks?id={id} - удаление задачи")

	log.Printf("Сервер запускается на порту %s", port)
	return http.ListenAndServe(":"+port, router)
}
