package main

import (
	"ToDoList/internal/models"
	"ToDoList/internal/services"
	"fmt"
)

func main() {
	us := services.UserService{}
	u := models.User{Username: "1", Email: "1"}

	for i := 0; i < 5; i++ {
		name := fmt.Sprint("Name", i)
		email := fmt.Sprint("Email", i)
		us.AddUser(name, "password", email)

	}
	us.Users[3] = u
	fmt.Printf("%#+v", us)
	fmt.Println("\n")
	for _, user := range us.Users {
		fmt.Printf("%#+v\n", user)
	}

	_, user, err := us.FindUserById(us.Users[0].Id)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#+v\n", user)

	err = us.UpdateUserById(user.Id, "", "", "newEmail")

	fmt.Printf("%#+v\n", us.Users[0])
	err = us.DeleteUserById(us.Users[3].Id)
	fmt.Println("\n")
	for _, user := range us.Users {
		fmt.Printf("%#+v\n", user)
	}
}
