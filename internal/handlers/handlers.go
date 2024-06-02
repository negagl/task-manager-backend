package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"task_manager_backend/internal/models"
	"task_manager_backend/internal/storage"
)

// HomeHandler Just says hi
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, filepath.Join("web", "index.html"))
}

// TasksHandler GET: Retrieves all tasks, POST: Post a new task
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var tasks []models.Task
		storage.DB.Find(&tasks)	// Find retrieves all records
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		var task models.Task
		json.NewDecoder(r.Body).Decode(&task)
		storage.DB.Create(&task)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// TaskHandler GET: Gets a task by its ID, PUT: Updates a task at a time by its ID, DELETE: Delete a task by its ID
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	id_txt := r.PathValue("id")
	id, err := strconv.Atoi(id_txt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var task models.Task

	switch r.Method {
	case http.MethodGet:
		// We check that record exists in DB. First searches by primary key (id)
		err := storage.DB.First(&task, id).Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(task)
	case http.MethodPut:
		err := storage.DB.First(&task, id).Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewDecoder(r.Body).Decode(&task)
		storage.DB.Save(&task)	// Function to update records (or create if not found)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	case http.MethodDelete:
		// Try to Delete if found a record. Else return error
		err := storage.DB.Delete(&task, id).Error
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
