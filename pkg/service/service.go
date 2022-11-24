package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateAccessToken(username, password string) (string, error)
	ParseAccessToken(token string) (int, error)
}

type ProductI interface {
	GetAll() ([]model.Product, error)
	GetAllCategories() ([]model.ProductCategory, error)
}

type Service struct {
	Authorization
	ProductI
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		ProductI:      NewProductService(repo.ProductI),
	}
}
