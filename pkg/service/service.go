package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateAccessToken(username, password string) (string, error)
	ParseAccessToken(token string) (int, error)
	GetUserById(id int) ([]model.User, error)
}
type UserCardI interface {
	GetAll(userId int) ([]model.UserCard, error)
	PostProductToCard(userId int, product model.Product) error
}

type ProductI interface {
	GetAll() ([]model.Product, error)
	GetAllCategories() ([]model.Categories, error)
	GetAllProviders() ([]model.Providers, error)
}

type Service struct {
	Authorization
	ProductI
	UserCardI
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		ProductI:      NewProductService(repo.ProductI),
		UserCardI:     NewUserCardService(repo.UserCardI),
	}
}
