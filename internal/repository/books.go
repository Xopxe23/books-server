package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/xopxe23/books-server/internal/domain"
)

type BooksPostgres struct {
	db *sqlx.DB
}

func NewBooksRepos(db *sqlx.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

func (b *BooksPostgres) Create(book domain.Book) (int, error) {
	row := b.db.QueryRow("INSERT INTO books (title, author) values ($1, $2) RETURNING id", book.Title, book.Author)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BooksPostgres) GetAll() ([]domain.Book, error) {
	var books []domain.Book
	if err := b.db.Select(&books, "SELECT * FROM books"); err != nil {
		return nil, err
	}
	return books, nil
}

func (b *BooksPostgres) GetById(id int) (domain.Book, error) {
	var book domain.Book
	if err := b.db.Get(&book, "SELECT * FROM books WHERE id = $1", id); err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (b *BooksPostgres) Update(id int, input domain.UpdateBookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, *input.Author)
		argId++
	}

	if input.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *input.Rating)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE books SET %s WHERE id = $%d", setQuery, argId)
	args = append(args, id)

	_, err := b.db.Exec(query, args...)
	return err
}

func (b *BooksPostgres) Delete(id int) error {
	_, err := b.db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
