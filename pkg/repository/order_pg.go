package repository

import (
	"context"
	"course_work/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderPostgres struct {
	db *pgxpool.Pool
}

func (o OrderPostgres) DeleteById(id int) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderPostgres) GetAll() ([]model.Order, error) {
	query := "SELECT orders.id, orders.date, orders.price, orders.cash, i.id,ii.itemname, c.id,c.firstname,e.id,e.firstname FROM orders\nJOIN employee e on e.id = orders.employee_id\nJOIN customer c on c.id = orders.customer_id\nJOIN item i on  i.id = orders.item_id\nJOIN item_info ii on ii.id = i.info_id"
	rows, err := o.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Order])
	return orders, err
}

func NewOrderPostgres(db *pgxpool.Pool) *OrderPostgres {
	return &OrderPostgres{db: db}
}
