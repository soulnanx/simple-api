package repositories

import (
	"simple-api/internal/domain"
	"sync"
)

type memoryTaskRepo struct {
	mu     sync.Mutex
	tasks  []domain.Task
	nextID int
}

func NewMemoryTaskRepo() domain.TaskRepository {
	return &memoryTaskRepo{
		tasks:  []domain.Task{},
		nextID: 1,
	}
}

func (r *memoryTaskRepo) GetAll() ([]domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.tasks, nil
}

func (r *memoryTaskRepo) Create(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, *task)
	return nil
}

func (r *memoryTaskRepo) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return domain.ErrTaskNotFound
}
