package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/VeeRomanoff/todo-app/domain"
	"github.com/VeeRomanoff/todo-app/internal/repository"
	"github.com/google/uuid"
)

const salt = "seilkdjhiu238hri8uh324f0hdu32hg4123"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user domain.User) (uuid.UUID, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
