package models

import "time"

type Task struct {
	Title       string
	Description string
	Status      StatusOfTask
	Priority    Priority
	Executor    User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
