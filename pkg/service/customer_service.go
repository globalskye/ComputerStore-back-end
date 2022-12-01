package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type CustomerService struct {
	repo repository.CustomerI
}

func (c CustomerService) GetAll() ([]model.Customer, error) {
	return c.repo.GetAll()
}

func NewCustomerService(repo repository.CustomerI) *CustomerService {
	return &CustomerService{repo: repo}
}
