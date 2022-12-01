package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type EmployeeService struct {
	repo repository.EmployeeI
}

func (e EmployeeService) GetAll() ([]model.Employee, error) {
	return e.repo.GetAll()
}

func NewEmployeeService(repo repository.EmployeeI) *EmployeeService {
	return &EmployeeService{repo: repo}
}
