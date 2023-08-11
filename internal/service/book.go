package service

import "github.com/xopxe23/books-server/internal/domain"

type BooksRepository interface {
	Create(book domain.Book) (int, error)
	GetAll() ([]domain.Book, error)
	GetById(id int) (domain.Book, error)
	Update(id int, input domain.UpdateBookInput) error
	Delete(id int) error
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

func (b *Books) GetById(id int) (domain.Book, error) {
	return b.repo.GetById(id)
}

func (b *Books) Update(id int, input domain.UpdateBookInput) error {
	return b.repo.Update(id, input)
}

func (b *Books) Delete(id int) error {
	return b.repo.Delete(id)
}
