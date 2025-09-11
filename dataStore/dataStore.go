package datastore

import (
	"sync"
	"tasks_manager/errors"
	"tasks_manager/task"
)

type DataStoreApi interface {
	CreateTask(task task.Task) error
	LoadTask(id int64) (task.Task, error)
	DeleteTask(id int64) error
	GetTasks() ([]task.Task, error)
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

func (ds *DataStore) CreateTask(t task.Task) error {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()

	if _, ok := ds.data[t.ID]; ok {
		return errors.ErrTaskAlreadyExists
	}

	ds.data[t.ID] = t
	return nil
}

func (ds *DataStore) LoadTask(id int64) (task.Task, error) {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	if t, ok := ds.data[id]; ok {
		return t, nil
	}
	return task.Task{}, errors.ErrTaskNotFound
}

func (ds *DataStore) DeleteTask(id int64) error {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	if _, ok := ds.data[id]; ok {
		delete(ds.data, id)
		return nil
	}
	return errors.ErrTaskNotFound
}

func (ds *DataStore) GetTasks() ([]task.Task, error) {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	tasks := make([]task.Task, 0, len(ds.data))
	for _, t := range ds.data {
		tasks = append(tasks, t)
	}
	return tasks, nil
}
