package domain

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `json:"uuid"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
