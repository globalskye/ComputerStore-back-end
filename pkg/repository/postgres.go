package repository

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"os"
)

func NewPostgresDb() (*pgx.Conn, error) {

	db, err := pgx.Connect(pgx.ConnConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetUint16("db.port"),
		Database: viper.GetString("db.dbname"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASS"),
	})
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil

}
