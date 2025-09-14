package datastore

import (
	"sync"
	"tasks_manager/errors"
	"tasks_manager/task"
)

type DataStoreApi interface {
	HandleCreateTask(task task.Task) error
	HandleUpdateTask(task task.Task) error
	HandleGetTask(id int64) (task.Task, error)
	HandleGetAllUncompletedTasks() []task.Task
	HandleDeleteTask(id int64) error
	HandleGetTasks() ([]task.Task, error)
}

type DataStore struct {
	data map[int64]task.Task
	mtx  sync.Mutex
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: make(map[int64]task.Task),
	}
}

func (ds *DataStore) HandleCreateTask(t task.Task) error {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()

	if _, ok := ds.data[t.ID]; ok {
		return errors.ErrTaskAlreadyExists
	}

	ds.data[t.ID] = t
	return nil
}

func (ds *DataStore) HandleUpdateTask(t task.Task) error {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()

	task, ok := ds.data[t.ID]
	if !ok {
		return errors.ErrTaskNotFound
	}

	task.Completed = t.Completed
	ds.data[t.ID] = task
	return nil
}

func (ds *DataStore) HandleGetTask(id int64) (task.Task, error) {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	if t, ok := ds.data[id]; ok {
		return t, nil
	}
	return task.Task{}, errors.ErrTaskNotFound
}

func (ds *DataStore) HandleGetAllUncompletedTasks() []task.Task {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()

	var result []task.Task
	for _, task := range ds.data {
		if !task.Completed {
			result = append(result, task)
		}
	}
	return result
}

func (ds *DataStore) HandleDeleteTask(id int64) error {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	if _, ok := ds.data[id]; ok {
		delete(ds.data, id)
		return nil
	}
	return errors.ErrTaskNotFound
}

func (ds *DataStore) HandleGetTasks() ([]task.Task, error) {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	tasks := make([]task.Task, 0, len(ds.data))
	for _, t := range ds.data {
		tasks = append(tasks, t)
	}
	return tasks, nil
}
