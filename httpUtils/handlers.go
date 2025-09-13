package httputils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	datastore "tasks_manager/dataStore"
	"tasks_manager/dto"
	apperrors "tasks_manager/errors"
	"tasks_manager/task"
	"time"

	"github.com/go-chi/chi/v5"
)

type HTTPHandlers struct {
	datastore datastore.DataStoreApi
}

func NewHTTPHandlers(ds datastore.DataStoreApi) *HTTPHandlers {
	return &HTTPHandlers{datastore: ds}
}

/*
pattern: /tasks
method:  POST
info:    JSON in HTTP request body

succeed:
  - status code:   201 Created
  - response body: JSON represent created task

failed:
  - status code:   400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto dto.TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errDTO := dto.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := taskDto.ValidateForCreate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errDTO := dto.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	createdTask := task.NewTask(taskDto.Title, taskDto.Description)

	if err := h.datastore.HandleCreateTask(*createdTask); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errDTO := dto.ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, apperrors.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(createdTask, "", "  ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response:", err)
		return
	}
}

/*
pattern: /tasks/{id}
method:  GET
info:    pattern

succeed:
  - status code: 200 Ok
  - response body: JSON represented found task

failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	stringId := chi.URLParam(r, "id")
	if id, err := strconv.ParseInt(stringId, 10, 64); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errDTO := dto.ErrorDTO{
			Message: "invalid task ID",
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	} else {
		if t, err := h.datastore.HandleGetTask(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errDTO := dto.ErrorDTO{
				Message: err.Error(),
				Time:    time.Now(),
			}
			if errors.Is(err, apperrors.ErrTaskNotFound) {
				http.Error(w, errDTO.ToString(), http.StatusNotFound)
			} else {
				http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			}
			return
		} else {
			if byte, err := json.Marshal(t); err != nil {
				panic(err)
			} else {
				w.WriteHeader(http.StatusOK)
				if _, err := w.Write(byte); err != nil {
					fmt.Println("failed to write response:", err)
					return
				}
			}
		}
	}
}

/*
pattern: /tasks
method:  GET
info:    -

succeed:
  - status code: 200 Ok
  - response body: JSON represented found tasks

failed:
  - status code: 400, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {}

/*
pattern: /tasks?completed=true <<--- ребята тут я зафакапил, конечно же, если мы получаем список НЕвыполненных задач, то в query параметре должно быть completed=false, а не true
method:  GET
info:    query params

succeed:
  - status code: 200 Ok
  - response body: JSON represented found tasks

failed:
  - status code: 400, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {}

/*
pattern: /tasks/{title}
method:  PATCH
info:    pattern + JSON in request body

succeed:
  - status code: 200 Ok
  - response body: JSON represented changed task

failed:
  - status code: 400, 409, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {}

/*
pattern: /tasks/{title}
method:  DELETE
info:    pattern

succeed:
  - status code: 204 No Content
  - response body: -

failed:
  - status code: 400, 404, 500, ...
  - response body: JSON with error + time
*/
func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {}
