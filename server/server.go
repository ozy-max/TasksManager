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

	// Добавляем middleware
	router.Use(Logger)
	router.Use(Recoverer)
	router.Use(RequestID)

	// Регистрируем маршруты

	// Debug: добавляем логирование всех маршрутов
	log.Println("Зарегистрированные маршруты:")
	log.Println("POST /tasks - создание задачи")
	log.Println("GET /tasks - получение всех задач")
	log.Println("GET /tasks/{id} - получение задачи по ID")
	log.Println("PATCH /tasks/{id} - обновление задачи")
	log.Println("DELETE /task?id={id} - удаление задачи")

	log.Printf("Сервер запускается на порту %s", port)
	return http.ListenAndServe(":"+port, router)
}
