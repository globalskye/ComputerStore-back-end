package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type UserCardService struct {
	repo repository.UserCardI
}

func (u UserCardService) GetAll(userId int) ([]model.UserCard, error) {
	//TODO implement me
	return u.repo.GetAll(userId)
}

func (u UserCardService) PostProductToCard(userId int, product model.Product) error {
	//TODO implement me
	return u.repo.PostProductToCard(userId, product)
}

func NewUserCardService(repo repository.UserCardI) *UserCardService {
	return &UserCardService{repo: repo}
}
