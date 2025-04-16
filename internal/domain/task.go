package domain

import "errors"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskRepository interface {
	GetAll() ([]Task, error)
	Create(task *Task) error
	Delete(id int) error
}
