package task

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewTask(title, description string) *Task {
	return &Task{
		ID:          int64(uuid.New().ID()),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()
	t.Completed = true
	t.UpdatedAt = &completeTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.UpdatedAt = nil
}
