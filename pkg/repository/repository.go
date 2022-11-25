package repository

import (
	"course_work/pkg/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type ProductI interface {
	GetAll() ([]model.Product, error)
	GetAllCategories() ([]model.Categories, error)
}

type Repository struct {
	Authorization
	ProductI
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ProductI:      NewProductPostgres(db),
	}
}
