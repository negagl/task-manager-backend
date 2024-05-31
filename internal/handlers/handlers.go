package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task_manager_backend/internal/models"
	"task_manager_backend/internal/storage"
)

// HomeHandler Just says hi
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Task Manager API"))
}

// TasksHandler GET: Retrieves all tasks, POST: Post a new task
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks := storage.GetTasks()
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		var task models.Task
		json.NewDecoder(r.Body).Decode(&task)
		storage.AddTask(task)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// TaskHandler GET: Gets a task by its ID, PUT: Updates a task at a time by its ID, DELETE: Delete a task by its ID
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tasks := storage.GetTasks()
		for _, task := range tasks {
			if task.ID == id {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(task)
				return
			}
		}
	case http.MethodPut:
		var updatedTask models.Task
		json.NewDecoder(r.Body).Decode(&updatedTask)
		storage.UpdateTask(id, updatedTask)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedTask)
	case http.MethodDelete:
		storage.DeleteTask(id)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
