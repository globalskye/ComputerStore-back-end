package repository

import (
	"course_work/pkg/model"
	"github.com/jackc/pgx"
)

type AuthPostgres struct {
	db *pgx.Conn
}

func NewAuthPostgres(db *pgx.Conn) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := "INSERT INTO users (username, email, password_hash) VALUES($1,$2,$3) RETURNING id"

	row := a.db.QueryRow(query, user.Username, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (a *AuthPostgres) GetUser(username, password string) (model.User, error) {
	query := "SELECT id FROM users WHERE username=$1 AND password_hash=$2"
	var user model.User
	row := a.db.QueryRow(query, username, password)
	err := row.Scan(&user.ID)

	return user, err
}
