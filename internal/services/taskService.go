package services

import (
	"ToDoList/internal/models"
	"time"
)

type TaskService struct {
	Task models.Task
}

func (ts *TaskService) CreateTask(
	title, description string,
	status models.StatusOfTask,
	priority models.Priority,
	executor *models.User) TaskService {
	timeCreated := time.Now()
	ts.Task = models.Task{
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
		Executor:    executor,
		CreatedAt:   timeCreated,
		UpdatedAt:   timeCreated,
	}
	return *ts
}

func (ts *TaskService) GetTask() models.Task {
	return ts.Task
}
