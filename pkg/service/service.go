package service

import "course_work/pkg/repository"

type Authorization interface {
}

type CourseWork interface {
}

type Service struct {
	Authorization
	CourseWork
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
