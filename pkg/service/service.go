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

type CustomerI interface {
	Create(customer model.Customer) (int, error)
	GetList() ([]model.Customer, error)
	GetById(id int) model.Customer
	Patch(customer model.Customer) (int, error)
	Delete(id int) error
}

type Service struct {
	Authorization
	CustomerI
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		CustomerI:     NewCustomerService(repo.CustomerI),
	}
}
