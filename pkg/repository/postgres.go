package repository

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"os"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

/*func NewPostgresDb(cfg DbConfig) (*sqlx.DB, error) {
	//postgres://username:password@localhost:5432/database_name
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLmode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}*/

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
