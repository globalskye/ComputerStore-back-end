package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type OrderService struct {
	repo repository.OrderI
}

func (o OrderService) CreateOrder(cards []model.UserCard, id int) error {
	return o.repo.CreateOrder(cards, id)
}

func (o OrderService) GetAll() ([]model.Order, error) {

	return o.repo.GetAll()
}

func NewOrderService(repo repository.OrderI) *OrderService {
	return &OrderService{repo: repo}
}
