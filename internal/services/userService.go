package services

import (
	"ToDoList/internal/models"
	"errors"
)

type UserService struct {
	Users []models.User
}

func (us *UserService) AddUser(name, password, email string) {
	modelUser := models.User{}
	user := modelUser.NewUser(name, password, email)
	us.Users = append(us.Users, *user)
}

func (us *UserService) FindUserById(id string) (int, *models.User, error) {
	countUsers := len(us.Users)
	for index := 0; index < countUsers; index++ {
		if us.Users[index].Id == id {
			return index, &us.Users[index], nil
		}
	}
	return -1, &models.User{}, errors.New("user с тамим id не существует")
}

func (us *UserService) UpdateUserById(id, name, password, email string) error {
	_, user, err := us.FindUserById(id)
	if err != nil {
		return err
	}
	user.UpdateUser(name, password, email)
	return nil
}

func (us *UserService) DeleteUserById(id string) error {
	index, _, err := us.FindUserById(id)
	if err != nil {
		return err
	}
	us.Users = append(us.Users[:index], us.Users[index+1:]...)
	return nil
}

func (us *UserService) GetAllUsers() []models.User {
	return us.Users
}
