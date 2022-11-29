package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func NewPostgresDb() (*pgxpool.Pool, error) {

	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}
