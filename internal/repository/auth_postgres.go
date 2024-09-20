package repository

import (
	"github.com/VeeRomanoff/todo-app/domain"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) CreateUser(user domain.User) (uuid.UUID, error) {
	return uuid.New(), nil
}
