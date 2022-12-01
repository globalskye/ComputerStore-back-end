package repository

import (
	"context"
	"course_work/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerPostgres struct {
	db *pgxpool.Pool
}

func (c CustomerPostgres) GetAll() ([]model.Customer, error) {
	query := "SELECT * FROM customer"
	rows, err := c.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	customers, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Customer])
	return customers, err
}

func NewCustomerPostgres(db *pgxpool.Pool) *CustomerPostgres {
	return &CustomerPostgres{db: db}
}
