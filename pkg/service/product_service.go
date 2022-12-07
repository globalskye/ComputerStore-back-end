package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type ProductService struct {
	repo repository.ProductI
}

func (p ProductService) PostProductToStock(product model.Stock) error {

	return p.repo.PostProductToStock(product)
}

func (p ProductService) DeleteById(id int) error {
	return p.repo.DeleteById(id)
}

func (p ProductService) GetAllProviders() ([]model.Providers, error) {

	return p.repo.GetAllProviders()
}

func (p ProductService) GetAllCategories() ([]model.Categories, error) {
	return p.repo.GetAllCategories()
}

func (p ProductService) GetAll() ([]model.Stock, error) {
	return p.repo.GetAll()

}

func NewProductService(repo repository.ProductI) *ProductService {
	return &ProductService{repo: repo}
}
