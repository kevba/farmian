package models

import "github.com/google/uuid"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(username, password string) *User {
	return &User{
		ID:       uuid.New().String(),
		Username: username,
		Password: password,
	}
}
