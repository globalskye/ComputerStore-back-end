package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateAccessToken(username, password string) (string, error)
}

type CourseWork interface {
}

type Service struct {
	Authorization
	CourseWork
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
