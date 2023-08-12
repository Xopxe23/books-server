package service

import (
	"github.com/xopxe23/books-server/internal/domain"
	"github.com/xopxe23/books-server/internal/repository"
)

type BooksService struct {
	repo repository.BooksRepository
}

func NewBooksService(repo repository.BooksRepository) *BooksService {
	return &BooksService{repo: repo}
}

func (b *BooksService) Create(book domain.Book) (int, error) {
	return b.repo.Create(book)
}

func (b *BooksService) GetAll() ([]domain.Book, error) {
	return b.repo.GetAll()
}

func (b *BooksService) GetById(id int) (domain.Book, error) {
	return b.repo.GetById(id)
}

func (b *BooksService) Update(id int, input domain.UpdateBookInput) error {
	return b.repo.Update(id, input)
}

func (b *BooksService) Delete(id int) error {
	return b.repo.Delete(id)
}
