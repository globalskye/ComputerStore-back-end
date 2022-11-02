package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
	"crypto/sha256"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}
type jwtTokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
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

func (a *AuthService) GenerateAccessToken(username, password string) (string, error) {
	user, err := a.repo.GetUser(username, generateHashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwtTokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		UserId: user.ID,
	})
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))

}
