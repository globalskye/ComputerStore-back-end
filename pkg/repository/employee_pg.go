package repository

import (
	"context"
	"course_work/pkg/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeePostgres struct {
	db *pgxpool.Pool
}

func (e EmployeePostgres) GetAll() ([]model.Employee, error) {
	query := "SELECT * FROM employee"
	rows, err := e.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	employee, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.Employee])
	return employee, err
}

func NewEmployeePostgres(db *pgxpool.Pool) *EmployeePostgres {
	return &EmployeePostgres{db: db}
}
