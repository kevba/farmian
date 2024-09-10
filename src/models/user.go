package models

import "github.com/google/uuid"

type User struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(username, password string) *User {
	return &User{
		ID:       uuid.New().ID(),
		Username: username,
		Password: password,
	}
}
