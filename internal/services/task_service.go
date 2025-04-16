package services

import (
	"simple-api/internal/domain"
)

type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]domain.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) CreateTask(task *domain.Task) error {
	task.IsCompleted = false
	return s.repo.Create(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
