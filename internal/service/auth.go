package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fanfaronDo/code_education_api/internal/domain"
	"github.com/fanfaronDo/code_education_api/internal/repository"
	"time"
)

const (
	salt      = "LWajdnwHFyglfQWFbi"
	signedKey = "HAUWFGiwgbkwsGeHGeuh"
	tokenTTL  = 12 * time.Hour
)

type claims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type Authorization struct {
	repo *repository.Repository
}

func NewAuthorization(repo *repository.Repository) *Authorization {
	return &Authorization{
		repo: repo,
	}
}

func (a *Authorization) CreateUser(user domain.User) (int, error) {
	user.Password = a.generatePasswordHash(user.Password)
	return a.repo.AuthRepository.CreateUser(user)
}

func (a *Authorization) GenerateToken(username, password string) (string, error) {
	user, err := a.repo.AuthRepository.GetUser(username, a.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id})

	return token.SignedString([]byte(signedKey))
}

func (a *Authorization) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signedKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*claims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	return claims.UserId, nil
}

func (a *Authorization) generatePasswordHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
