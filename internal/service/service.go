package service

import (
	"github.com/VeeRomanoff/todo-app/domain"
	"github.com/VeeRomanoff/todo-app/internal/repository"
	"github.com/google/uuid"
)

type IAuthorizationService interface {
	CreateUser(user domain.User) (uuid.UUID, error)
}

type ITodoListService interface {
}

type ITodoItemService interface {
}

type Service struct {
	IAuthorizationService
	ITodoListService
	ITodoItemService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ITodoItemService: NewAuthService(repos.Authorization),
	}
}
