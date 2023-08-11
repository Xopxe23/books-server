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
