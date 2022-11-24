package repository

import (
	"course_work/pkg/model"
	"github.com/jackc/pgx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type ProductI interface {
	GetAll() ([]model.Product, error)
}

type Repository struct {
	Authorization
	ProductI
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ProductI:      NewProductPostgres(db),
	}
}
