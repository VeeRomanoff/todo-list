package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UUID      uuid.UUID `json:"-"`
	Email     string    `json:"email" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewUser(email string, username string, password string) *User {
	return &User{
		UUID:      uuid.New(),
		Email:     email,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
