package main

import (
	"log"
	"net/http"

	"github.com/AminMousaviUnity/ToGo/config"
	"github.com/AminMousaviUnity/ToGo/internal/api"
	"github.com/AminMousaviUnity/ToGo/internal/db"
	"github.com/AminMousaviUnity/ToGo/internal/repository"
	"github.com/AminMousaviUnity/ToGo/internal/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configs.")
	}

	// Run db migrations
	db.RunMigrations()

	// Connect to the database
	database := config.ConnectDB()
	defer database.Close()

	// Init repo, service, and handler layers
	repo := &repository.TaskRepository{DB: database}
	service := &service.TaskService{Repo: repo}
	handler := &api.Handler{Service: service}

	// Set up routes
	router := mux.NewRouter()
	router.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handler.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")

	// Start the server
	addr := ":6666"
	log.Printf("Server running on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
