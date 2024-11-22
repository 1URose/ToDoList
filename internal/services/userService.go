package services

import "ToDoList/internal/models"

type UserService struct{}

func (us *UserService) CreateUser(name, password, email string) *models.User {
	user := models.User{
		Username: name,
		Password: password,
		Email:    email,
	}
	return &user
}

func (us *UserService) GetUserName(user models.User) string {
	return user.Username
}

func (us *UserService) GetUserPassword(user models.User) string {
	return user.Password

}

func (us *UserService) GetUserEmail(user models.User) string {
	return user.Email
}

func (us *UserService) SetUserName(user *models.User, newName string) {
	user.Username = newName
}
func (us *UserService) SetUserPassword(user *models.User, newPassword string) {
	user.Password = newPassword
}
func (us *UserService) SetUserEmail(user *models.User, newEmail string) {
	user.Email = newEmail
}

func (us *UserService) DeleteUser(user *models.User) {
	user = nil
}
