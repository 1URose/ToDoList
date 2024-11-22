package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id          string
	Title       string
	Description string
	Status      StatusOfTask
	Priority    Priority
	Executor    *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) NewTask(title, description string, status StatusOfTask, priority Priority, executor *User) *Task {
	id := uuid.New().String()
	task := Task{
		Id:          id,
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
		Executor:    executor,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return &task
}

func (t *Task) UpdateTask(title, description string, status StatusOfTask, priority Priority, executor *User) {
	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	if status != 0 {
		t.Status = status
	}
	if priority != 0 {
		t.Priority = priority
	}
	if executor != nil {
		t.Executor = executor
	}
	t.UpdatedAt = time.Now()
}
