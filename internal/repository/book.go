package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/xopxe23/books-server/internal/domain"
)

type Books struct {
	db *sqlx.DB
}

func NewBooks(db *sqlx.DB) *Books {
	return &Books{db: db}
}

func (b *Books) Create(book domain.Book) (int, error) {
	row := b.db.QueryRow("INSERT INTO books (title, author) values ($1, $2) RETURNING id", book.Title, book.Author)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (b *Books) GetAll() ([]domain.Book, error) {
	var books []domain.Book
	if err := b.db.Select(&books, "SELECT * FROM books"); err != nil {
		return nil, err
	}
	return books, nil
}

func (b *Books) GetById(id int) (domain.Book, error) {
	var book domain.Book
	if err := b.db.Get(&book, "SELECT * FROM books WHERE id = $1", id); err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (b *Books) Delete(id int) error {
	_, err := b.db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
