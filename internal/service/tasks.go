package service

import (
	"errors"
	"github.com/AminMousaviUnity/ToGo/internal/models"
	"github.com/AminMousaviUnity/ToGo/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

// CreateTask validates and creates a new task
func (s *TaskService) CreateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title cannot be empty")
	}

	if task.Status == "" {
		task.Status = "pending"
	}

	return s.Repo.CreateTask(task)
}

// GetTask retrieves a single task by its ID
func (s *TaskService) GetTask(id int) (*models.Task, error) {
	task, err := s.Repo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetAllTasks retrieves all tasks
func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	tasks, err := s.Repo.GetAllTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// UpdateTask updates an existing task by ID
func (s *TaskService) UpdateTask(id int, updates *models.Task) error {
	if updates.Title == "" {
		return errors.New("title cannot be empty")
	}
	
	task, err := s.Repo.GetTaskByID(id)
	if err != nil {
		return err
	}

	if updates.Description != "" {
		task.Description = updates.Description
	}
	if updates.Status != "" {
		task.Status = updates.Status
	}
	if !updates.DueDate.IsZero() {
		task.DueDate = updates.DueDate
	}

	return s.Repo.UpdateTask(task)
}

// DeleteTask deletes a task by its ID
func (s *TaskService) DeleteTask(id int) error {
	return s.Repo.DeleteTask(id)
}
