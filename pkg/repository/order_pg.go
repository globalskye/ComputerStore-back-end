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

func (o OrderPostgres) CreateOrder(card model.UserCard) error {
	query := "UPDATE users SET adress=$1 WHERE id=$2"
	_, err := o.db.Query(context.Background(), query, card.UserAdress, card.UserId)
	if err != nil {
		return err
	}
	for _, v := range card.Items {

		query := "UPDATE mainstock SET itemcount=itemcount-$1 WHERE item_id=$2"
		_, err := o.db.Query(context.Background(), query, v.Count, v.Item.Id)
		if err != nil {
			return err
		}
		query = "INSERT INTO orders(date, price, cash, taxes, item_id, item_count, user_id, employee_id, ksa_id)" +
			"VALUES(now(),$1,false,1,$2,$3,$4,1,1)"
		_, err = o.db.Query(context.Background(), query, v.TotalPrice, v.Item.Id, v.Count, card.UserId)
		if err != nil {
			return err
		}

		return err
	}

	return nil
}

func (o OrderPostgres) DeleteById(id int) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderPostgres) GetAll() ([]model.Order, error) {
	query := "SELECT orders.id, date,price,cash,item.id,item_count, item_info.itemname,users.id,users.username,e.id,e.firstName FROM orders\n    JOIN employee e on e.id = orders.employee_id\n    JOIN item on orders.item_id = item.id\n    JOIN item_info on item.info_id = item_info.id\n    JOIN users on orders.user_id = users.id\n"
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
