package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type CustomerService struct {
	repo repository.CustomerI
}

func (c CustomerService) Create(customer model.Customer) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerService) GetList() ([]model.Customer, error) {

	panic("implement me")
}

func (c CustomerService) GetById(id int) model.Customer {
	//TODO implement me
	panic("implement me")
}

func (c CustomerService) Patch(customer model.Customer) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerService) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewCustomerService(repo repository.CustomerI) *CustomerService {
	return &CustomerService{repo: repo}
}
