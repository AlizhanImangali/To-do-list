package service

import (
	"TODO"
	"TODO/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "sdjadja51da5dda"

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) GenerateToken(username, password string) (string error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user TODO.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
