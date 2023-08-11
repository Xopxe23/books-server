package service

import "github.com/xopxe23/books-server/internal/domain"

type BooksRepository interface {
	Create(book domain.Book) (int, error)
	GetAll() ([]domain.Book, error)
}

type Books struct {
	repo BooksRepository
}

func NewBooks(repo BooksRepository) *Books {
	return &Books{
		repo: repo,
	}
}

func (b *Books) Create(book domain.Book) (int, error) {
	return b.repo.Create(book)
}

func (b *Books) GetAll() ([]domain.Book, error) {
	return b.repo.GetAll()
}
