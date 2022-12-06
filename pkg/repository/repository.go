package repository

import (
	"course_work/pkg/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
	GetUserById(id int) ([]model.User, error)
}
type UserCardI interface {
	GetAll(userId int) ([]model.UserCard, error)
	PostProductToCard(userId int, product model.Product) error
}
type CustomerI interface {
	GetAll() ([]model.Customer, error)
}
type EmployeeI interface {
	GetAll() ([]model.Employee, error)
}

type ProductI interface {
	GetAll() ([]model.Product, error)
	GetAllCategories() ([]model.Categories, error)
	GetAllProviders() ([]model.Providers, error)
}
type UserI interface {
	GetAll() ([]model.User, error)
}

type Repository struct {
	Authorization
	ProductI
	UserCardI
	CustomerI
	EmployeeI
	UserI
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ProductI:      NewProductPostgres(db),
		UserCardI:     NewCardPostgres(db),
		CustomerI:     NewCustomerPostgres(db),
		EmployeeI:     NewEmployeePostgres(db),
		UserI:         NewUserPostgres(db),
	}
}
