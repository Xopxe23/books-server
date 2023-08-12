package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/xopxe23/books-server/internal/domain"
	"github.com/xopxe23/books-server/internal/repository"
)

const (
	solt = "fdagjgndnafjff"
)

type AuthService struct {
	repo repository.UsersRepository
}

func NewAuthService(repo repository.UsersRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(input domain.User) (int, error) {
	input.Password = generatePasswordHash(input.Password)
	return s.repo.CreateUser(input)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}
