package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/xopxe23/books-server/internal/domain"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthRepos(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (p *AuthPostgres) CreateUser(input domain.User) (int, error) {
	var id int
	row := p.db.QueryRow("INSERT INTO users(name, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
		input.Name, input.Email, input.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
