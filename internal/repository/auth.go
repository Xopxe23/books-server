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

func (p *AuthPostgres) CreateUser(input domain.SignUpInput) (int, error) {
	var id int
	row := p.db.QueryRow("INSERT INTO users(name, email, password_hash) VALUES ($1, $2, $3) RETURNING id",
		input.Name, input.Email, input.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *AuthPostgres) GetUser(input domain.SignInInput) (domain.User, error) {
	var user domain.User
	err := p.db.Get(&user, "SELECT * FROM users WHERE email = $1 and password_hash = $2",
		input.Email, input.Password)
	
	return user, err
}
