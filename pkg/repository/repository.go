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

type ProductI interface {
	GetAll() ([]model.Product, error)
	GetAllCategories() ([]model.Categories, error)
	GetAllProviders() ([]model.Providers, error)
}

type Repository struct {
	Authorization
	ProductI
	UserCardI
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		ProductI:      NewProductPostgres(db),
		UserCardI:     NewCardPostgres(db),
	}
}
