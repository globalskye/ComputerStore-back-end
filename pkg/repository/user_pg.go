package repository

import (
	"context"
	"course_work/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgres struct {
	db *pgxpool.Pool
}

func (u UserPostgres) DeleteById(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := u.db.Query(context.TODO(), query, id)
	return err
}

func (u UserPostgres) GetAll() ([]model.User, error) {
	query := "SELECT * FROM users"
	rows, err := u.db.Query(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.User])
	return users, err
}

func NewUserPostgres(db *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{db: db}
}
