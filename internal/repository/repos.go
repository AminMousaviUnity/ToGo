package repository

import (
	"github.com/AminMousaviUnity/ToGo/internal/models"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	DB *sqlx.DB
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := `INSERT INTO tasks (title, description, status, due_data)
			  VALIES ($1, $2, $3, $4) RETURNING id`
	
	return r.DB.QueryRow(query, task.Title, task.Description, task.DueDate).Scan(&task.ID)
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Select(&tasks, "SELECT * FROM ORDER BY created_at DESC")
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	query := "SELECT * FROM tasks WHERE id = $1"
	err := r.DB.Get(&task, query, id)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2, status = $3, due_date = $4, updated_at = NOW() WHERE id = $5`
	_, err := r.DB.Exec(query, task.Title, task.Description, task.Status, task.DueDate, task.ID)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
