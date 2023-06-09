package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/antonchaban/file-hub-go/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "yfbhHFDCJHKAJH6546"
	signingKey = "ssefyuhfnhisdhj25345"
	tokenTTL   = time.Hour * 24
)

type Authorization interface {
	CreateUser(user fhub.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	InvalidateToken(accessToken string) (int, error)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {

	isBlacklisted, err := s.repo.IsTokenInBlacklist(accessToken)
	fmt.Println("accestoken", accessToken)
	fmt.Println("isblacklisted", isBlacklisted)
	if err != nil {
		return 0, err
	}
	if isBlacklisted {
		return 0, errors.New("token invalidated")
	}

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo}
}

func (s *AuthService) CreateUser(user fhub.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) InvalidateToken(accessToken string) (int, error) {
	id, err := s.repo.AddTokenToBlacklist(accessToken)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
