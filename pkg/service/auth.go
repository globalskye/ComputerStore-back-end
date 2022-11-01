package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
	"crypto/sha256"
	"fmt"
	"os"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generateHashPassword(user.Password)
	return a.repo.CreateUser(user)
}
func generateHashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}
