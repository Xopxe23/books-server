package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
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

func (s *AuthService) CreateUser(input domain.SignUpInput) (int, error) {
	input.Password = generatePasswordHash(input.Password)
	return s.repo.CreateUser(input)
}

func (s *AuthService) GenerateToken(input domain.SignInInput) (string, error) {
	input.Password = generatePasswordHash(input.Password)
	user, err := s.repo.GetUser(input)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   strconv.Itoa(int(user.ID)),
		"iat":  time.Now().Unix(),
		"exp": time.Now().Add(tokenTTL).Unix(),
	})
	return token.SignedString([]byte(signingKey))
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
