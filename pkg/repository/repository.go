package repository

import (
	"github.com/jackc/pgx"
)

type Authorization interface {
}

type CourseWork interface {
}

type Repository struct {
	Authorization
	CourseWork
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{}
}
