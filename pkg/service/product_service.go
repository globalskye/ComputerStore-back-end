package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type ProductService struct {
	repo repository.ProductI
}

func (p ProductService) GetAllCategories() ([]model.ProductCategory, error) {
	return p.repo.GetAllCategories()
}

func (p ProductService) GetAll() ([]model.Product, error) {
	return p.repo.GetAll()

}

func NewProductService(repo repository.ProductI) *ProductService {
	return &ProductService{repo: repo}
}
