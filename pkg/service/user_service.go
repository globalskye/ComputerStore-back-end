package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type UserService struct {
	repo repository.UserI
}

func (u UserService) GetAll() ([]model.User, error) {

	return u.repo.GetAll()
}

func NewUserService(repo repository.UserI) *UserService {
	return &UserService{repo: repo}
}
