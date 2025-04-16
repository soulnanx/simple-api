package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	_ "simple-api/docs"
	"sync"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}

var (
	tasks  = []Task{}
	nextID = 1
	mu     sync.Mutex
)

// @title API de Tasks
// @version 1.0
// @description Esta API gerencia tarefas
// @host localhost:3000
// @BasePath /
func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	swaggerHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/ping", pingHandler)
	http.Handle("/swagger/", swaggerHandler)

	fmt.Printf("Servidor rodando na porta %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTasks(w, r)
	case "POST":
		addTask(w, r)
	case "DELETE":
		deleteTask(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	ping(w, r)
}

func ping(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "pong")
}

// @Summary Busca todas as tasks
// @Description busca todas as tasks cadastradas
// @Tags tasks
// @Produce json
// @Success 200 {array} Task "Lista de tasks"
// @Router /tasks [get]
func getTasks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newTask.ID = nextID
	nextID++
	newTask.IsCompleted = false
	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	taskID := r.URL.Query().Get("id")
	for i, task := range tasks {
		if task.Title == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}
