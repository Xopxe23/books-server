package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/xopxe23/books-server/internal/domain"
)

type BooksRepository interface {
	Create(book domain.Book) (int, error)
	GetAll() ([]domain.Book, error)
	GetById(id int) (domain.Book, error)
	Update(id int, input domain.UpdateBookInput) error
	Delete(id int) error
}

type UsersRepository interface {
	SignUp(input domain.User) (int, error)
}

type Repository struct {
	BooksRepository
	UsersRepository
}

func NewRepostory(db *sqlx.DB) *Repository {
	return &Repository{
		BooksRepository: NewBooksRepos(db),
		UsersRepository: NewAuthRepos(db),
	}
}
