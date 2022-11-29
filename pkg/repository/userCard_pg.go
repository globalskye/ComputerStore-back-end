package repository

import (
	"context"
	"course_work/pkg/model"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CardPostgres struct {
	db *pgxpool.Pool
}

func (c CardPostgres) GetAll(userId int) ([]model.UserCard, error) {
	query := "SELECT * FROM user_card WHERE user_id=$1"
	rows, err := c.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	card, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.UserCard])
	return card, err
}

func (c CardPostgres) PostProductToCard(userId int, product model.Product) error {
	query := "INSERT INTO user_card(user_id, item_id) VALUES($1,$2)"
	_, err := c.db.Query(context.Background(), query, userId, product.ID)
	if err != nil {
		return err
	}
	fmt.Println(err)
	return err
}

func NewCardPostgres(db *pgxpool.Pool) *CardPostgres {
	return &CardPostgres{db: db}
}
