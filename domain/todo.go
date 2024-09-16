package domain

import "github.com/google/uuid"

type TodoList struct {
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type UsersList struct {
	UUID     uuid.UUID `json:"uuid"`
	UserUUID uuid.UUID `json:"userUUID"`
	ListUUID uuid.UUID `json:"listUUID"`
}

type TodoItem struct {
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

type ListItem struct {
	UUID     uuid.UUID `json:"uuid"`
	ListUUID uuid.UUID `json:"listUUID"`
	ItemUUID uuid.UUID `json:"itemUUID"`
}
