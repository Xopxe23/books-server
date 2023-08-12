package service

import (
	"github.com/xopxe23/books-server/internal/domain"
	"github.com/xopxe23/books-server/internal/repository"
)

type AuthService struct {
	repo repository.UsersRepository
}

func NewAuthService(repo repository.UsersRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(input domain.User) (int, error) {
	return s.repo.SignUp(input)
}
