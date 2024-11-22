package main

import (
	"ToDoList/internal/models"
	"fmt"
)

func main() {
	// Создаем объект User
	userr := models.User{}
	taskk := models.Task{}
	//Создаем объект Task
	user := userr.NewUser("User 1", "Password 1", "Email 1")

	newTask := taskk.NewTask(
		"Task 1",
		"Description 1",
		models.NeedToDo,
		models.High,
		user,
	)

	fmt.Printf("%#+v\n", *newTask)
	fmt.Printf("%#+v\n", (*newTask).Executor)
	user.Username = "User 2"
	fmt.Println()
	fmt.Printf("%#+v\n", *newTask)
	fmt.Printf("%#+v\n", (*newTask).Executor)
}
