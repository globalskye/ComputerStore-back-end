package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
	"fmt"
)

type OrderService struct {
	repo repository.OrderI
}

func (o OrderService) CreateOrder(card model.UserCard) error {
	newCard := card

	for i, v := range card.Items {
		newCard.Items[i].TotalPrice = v.Count * v.Item.Price
	}
	fmt.Println(newCard)
	return o.repo.CreateOrder(newCard)
}

func (o OrderService) GetAll() ([]model.Order, error) {

	return o.repo.GetAll()
}

func NewOrderService(repo repository.OrderI) *OrderService {
	return &OrderService{repo: repo}
}
