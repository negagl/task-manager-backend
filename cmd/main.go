package main

import (
	"log"
	"net/http"
	"task_manager_backend/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/tasks", handlers.TasksHandler)
	mux.HandleFunc("/task/", handlers.TaskHandler)

	log.Println("Server Started at :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
