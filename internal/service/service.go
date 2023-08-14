package service

import (
	"github.com/xopxe23/books-server/internal/domain"
	"github.com/xopxe23/books-server/internal/repository"
)

type Books interface {
	Create(book domain.Book) (int, error)
	GetAll() ([]domain.Book, error)
	GetById(id int) (domain.Book, error)
	Update(id int, input domain.UpdateBookInput) error
	Delete(id int) error
}

type Auth interface {
	SignUp(input domain.SignUpInput) (int, error)
	SignIn(input domain.SignInInput) (string, string, error)
	ParseToken(token string) (int64, error)
	RefreshTokens(refreshToken string) (string, string, error)
}

type Service struct {
	Books
	Auth
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books: NewBooksService(repos.BooksRepository),
		Auth:  NewAuthService(repos.UsersRepository),
	}
}
