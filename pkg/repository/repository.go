package repository

import (
	"course_work/pkg/model"
	"github.com/jackc/pgx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type CustomerI interface {
	Create(customer model.Customer) (int, error)
	GetList() ([]model.Customer, error)
	GetById(id int) model.Customer
	Patch(customer model.Customer) (int, error)
	Delete(id int) error
}

type Repository struct {
	Authorization
	CustomerI
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		CustomerI:     NewCustomerPostgres(db),
	}
}
