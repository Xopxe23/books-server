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

func (a *AuthPostgres) SignUp(input domain.User) (int, error) {
	return 0, nil
}
