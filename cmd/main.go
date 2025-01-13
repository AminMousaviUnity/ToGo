package main

import (
	"log"
	"net/http"

	"github.com/AminMousaviUnity/ToGo/config"
	"github.com/AminMousaviUnity/ToGo/internal/api"
	"github.com/AminMousaviUnity/ToGo/internal/db"
	"github.com/AminMousaviUnity/ToGo/internal/repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configuations.")
	}

	db.RunMigrations()


	db := config.ConnectDB()
	defer db.Close()

	repo := &repository.TaskRepository{DB: db}
	handler := &api.Handler{Repo: repo}

	router := mux.NewRouter()
	router.HandleFunc("/tasks", handler.GetTasks).Methods("GET")

	Addr := ":6666"
	log.Printf("Server running on port %s", Addr)
	log.Fatal(http.ListenAndServe(Addr, nil))
}
