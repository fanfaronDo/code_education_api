package service

import (
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"github.com/fanfaronDo/code_education_api/internal/repository"
)

type AuthService interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username, password string) (domain.User, error)
}

type Service struct {
	AuthService
}

func NewAuthService(repo repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo),
	}
}
