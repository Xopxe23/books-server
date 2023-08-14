package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/xopxe23/books-server/internal/domain"
	"github.com/xopxe23/books-server/internal/repository"
)

const (
	solt       = "fdagjgndnafjff"
	signingKey = "fdanfjkdafanfa"
	tokenTTL   = time.Hour * 12
)

type AuthService struct {
	repo repository.UsersRepository
}

func NewAuthService(repo repository.UsersRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(input domain.SignUpInput) (int, error) {
	input.Password = generatePasswordHash(input.Password)
	return s.repo.CreateUser(input)
}

func (s *AuthService) SignIn(input domain.SignInInput) (string, string, error) {
	input.Password = generatePasswordHash(input.Password)
	user, err := s.repo.GetUser(input)
	if err != nil {
		return "", "", err
	}
	return s.generateTokens(user.ID)
}

func (s *AuthService) generateTokens(id int64) (string, string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(id)),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(tokenTTL).Unix(),
	})
	assessToken, err := t.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := newRefreshToken()
	if err != nil {
		return "", "", err
	}

	if err := s.repo.CreateToken(domain.RefreshSession{
		UserID:    id,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}); err != nil {
		return "", "", err
	}

	return assessToken, refreshToken, nil
}

func newRefreshToken() (string, error) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func (s *AuthService) ParseToken(token string) (int64, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	fmt.Println(t.Claims)
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("invalid subject")
	}
	id, err := strconv.Atoi(sub)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}

func (s *AuthService) RefreshTokens(refreshToken string) (string, string, error) {
	session, err := s.repo.GetToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	if session.ExpiresAt.Unix() < time.Now().Unix() {
		return "", "", errors.New("refresh token expired")
	}
	return s.generateTokens(session.UserID)
}
