package datastore

import (
	"sync"
	"tasks_manager/httputils"
	"tasks_manager/task"
)

type DataStore interface {
	Save(key string, value []byte) error
	LoadTask(key string) ([]byte, error)
	DeleteTask(key string) error
	GetTasks() (map[string][]byte, error)
}

type DataStoreImpl struct {
	data map[int64]task.Task
	mtx  sync.Mutex
}

func NewDataStore() *DataStoreImpl {
	return &DataStoreImpl{
		data: make(map[int64]task.Task),
	}
}

func (ds *DataStoreImpl) Save(t task.Task) error {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()

	if _, ok := ds.data[t.ID]; ok {
		return httputils.ErrTaskAlreadyExists
	}

	ds.data[t.ID] = t
	return nil
}

func (ds *DataStoreImpl) LoadTask(id int64) (task.Task, error) {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	if t, ok := ds.data[id]; ok {
		return t, nil
	}
	return task.Task{}, httputils.ErrTaskNotFound
}

func (ds *DataStoreImpl) DeleteTask(id int64) error {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	if _, ok := ds.data[id]; ok {
		delete(ds.data, id)
		return nil
	}
	return httputils.ErrTaskNotFound
}

func (ds *DataStoreImpl) GetTasks() ([]task.Task, error) {
	ds.mtx.Lock()
	defer ds.mtx.Unlock()
	tasks := make([]task.Task, 0, len(ds.data))
	for _, t := range ds.data {
		tasks = append(tasks, t)
	}
	return tasks, nil
}
