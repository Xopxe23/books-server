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

func (p *AuthPostgres) CreateToken(token domain.RefreshSession) error {
	_, err := p.db.Exec("INSERT INTO refresh_tokens (user_id, token, expires_at) values ($1, $2, $3)",
		token.UserID, token.Token, token.ExpiresAt)
	return err
}

func (p *AuthPostgres) GetToken(token string) (domain.RefreshSession, error) {
	var t domain.RefreshSession
	err := p.db.QueryRow("SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token=$1", token).
		Scan(&t.ID, &t.UserID, &t.Token, &t.ExpiresAt)
	if err != nil {
		return t, err
	}

	_, err = p.db.Exec("DELETE FROM refresh_tokens WHERE user_id=$1", t.UserID)

	return t, err
}
