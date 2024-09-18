package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// поскольку репозитории работают с бд, передаем ему db *sqlx.DB
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
