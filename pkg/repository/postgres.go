package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
	"os"
)

func NewPostgresDb() (*pgx.Conn, error) {

	config := pgx.ConnConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetUint16("db.port"),
		Database: viper.GetString("db.dbname"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASS"),
	}
	db, err := pgx.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil

}
