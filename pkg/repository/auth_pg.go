package repository

import (
	"context"
	"course_work/pkg/model"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthPostgres struct {
	db *pgxpool.Pool
}

func NewAuthPostgres(db *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{db: db}
}

type exist struct {
	UserExist bool `json:"aa"`
}

func (a *AuthPostgres) CreateUser(user model.User) (int, error) {

	var id int
	query := "INSERT INTO users (username, email, password_hash) VALUES($1,$2,$3) RETURNING id"

	row := a.db.QueryRow(context.Background(), query, user.Username, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, errors.New("User with that username or email alredy exist")
	}
	return id, nil
}
func (a *AuthPostgres) GetUser(username, password string) (model.User, error) {

	query := "SELECT id FROM users WHERE username=$1 AND password_hash=$2"
	var user model.User
	row := a.db.QueryRow(context.Background(), query, username, password)
	err := row.Scan(&user.ID)

	return user, err
}
