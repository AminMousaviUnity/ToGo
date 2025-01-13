package repository

import (
	"github.com/AminMousaviUnity/ToGo/internal/models"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	DB *sqlx.DB
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Select(&tasks, "SELECT * FROM ORDER BY created_at DESC")
	return tasks, err
}
