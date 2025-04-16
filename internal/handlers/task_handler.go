package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"simple-api/internal/domain"
	"simple-api/internal/repositories"
	"simple-api/internal/services"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler() *TaskHandler {
	repo := repositories.NewMemoryTaskRepo()
	service := services.NewTaskService(repo)
	return &TaskHandler{service: service}
}

// @Summary Busca todas as tasks
// @Description busca todas as tasks cadastradas
// @Tags tasks
// @Produce json
// @Success 200 {array} domain.Task "Lista de tasks"
// @Router /tasks [get]
func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTasks(w, r)
	case http.MethodPost:
		h.addTask(w, r)
	case http.MethodDelete:
		h.deleteTask(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func (h *TaskHandler) addTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(id); err != nil {
		if err == domain.ErrTaskNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
