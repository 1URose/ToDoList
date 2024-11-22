package models

import "github.com/google/uuid"

type User struct {
	Id       string
	Username string
	password string
	Email    string
}

func (u *User) NewUser(name, password, email string) *User {
	id := uuid.New()
	user := User{
		Id:       id.String(),
		Username: name,
		password: password,
		Email:    email,
	}
	return &user
}

func (u *User) GetUserPassword(user User) string {
	return user.password
}

func (u *User) UpdateUser(newName, newPassword, newEmail string) error {
	if newName != "" {
		u.Username = newName
	}
	if newPassword != "" {
		u.password = newPassword
	}
	if newEmail != "" {
		u.Email = newEmail
	}
	return nil
}

func (u *User) ValidatePassword(inputPassword string) bool {
	return u.password == inputPassword
}
