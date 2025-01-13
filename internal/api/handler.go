package api

import (
	"encoding/json"
	"net/http"

	"github.com/AminMousaviUnity/ToGo/internal/repository"
)

type Handler struct {
	Repo *repository.TaskRepository
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.Repo.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}
