package services

import (
	"ToDoList/internal/models"
	"errors"
)

type TaskService struct {
	Task []models.Task
}

func (ts *TaskService) addTask(
	title, description string,
	status models.StatusOfTask,
	priority models.Priority,
	executor *models.User) {
	modelTask := models.Task{}
	newTask := modelTask.NewTask(title, description, status, priority, executor)
	ts.Task = append(ts.Task, *newTask)
}

func (ts *TaskService) GetTaskById(id string) (int, *models.Task, error) {
	countTasks := len(ts.Task)
	for index := 0; index < countTasks; index++ {
		if ts.Task[index].Id == id {
			return index, &ts.Task[index], nil
		}
	}
	return -1, &models.Task{}, errors.New("task с тамим id не существует")
}

func (ts *TaskService) UpdateTaskById(id string, title, description string, status models.StatusOfTask, priority models.Priority, executor *models.User) error {
	_, task, err := ts.GetTaskById(id)
	if err != nil {
		return err
	}
	task.UpdateTask(title, description, status, priority, executor)
	return nil
}

func (ts *TaskService) DeleteTaskById(id string) error {
	index, _, err := ts.GetTaskById(id)
	if err != nil {
		return err
	}
	ts.Task = append(ts.Task[:index], ts.Task[index+1:]...)
	return nil
}

func (ts *TaskService) GetAllTasks() []models.Task {
	return ts.Task
}
