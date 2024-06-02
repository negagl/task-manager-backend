package main

import (
	"log"
	"net/http"
	"task_manager_backend/internal/handlers"
	"task_manager_backend/internal/storage"
)

func main() {
	// Init DB
	storage.InitDB()

	// New Server
	mux := http.NewServeMux()

	// handle routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/tasks", handlers.TasksHandler)
	mux.HandleFunc("/task/{id}", handlers.TaskHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs) )

	// listen
	log.Println("Server Started at :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
