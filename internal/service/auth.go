package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fanfaronDo/to_do/internal/domain"
	"github.com/fanfaronDo/to_do/internal/repository"
	"time"
)

const (
	salt      = "DAWawdwassyglfQWFbi"
	signedKey = "DWAWDwassuh"
	tokenTTL  = 5 * time.Hour
)

type claims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type Authorization struct {
	repo *repository.Repository
}

func NewAuthorization(repo *repository.Repository) *Authorization {
	return &Authorization{repo}
}

func (r *Authorization) CreateUser(user domain.User) (int, error) {
	user.Password = r.generatePasswordHash(user.Password)
	return r.repo.AuthorisationRepository.CreateUser(user)
}

func (r *Authorization) GenerateToken(username, password string) (string, error) {
	user, err := r.repo.AuthorisationRepository.GetUser(username, r.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

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
