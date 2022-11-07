package repository

import (
	"course_work/pkg/model"
	"github.com/jackc/pgx"
)

type CustomerPostgres struct {
	db *pgx.Conn
}

func (c CustomerPostgres) Create(customer model.Customer) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerPostgres) GetList() ([]model.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerPostgres) GetById(id int) model.Customer {
	//TODO implement me
	panic("implement me")
}

func (c CustomerPostgres) Patch(customer model.Customer) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c CustomerPostgres) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewCustomerPostgres(db *pgx.Conn) *CustomerPostgres {
	return &CustomerPostgres{db: db}
}
