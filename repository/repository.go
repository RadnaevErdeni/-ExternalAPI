package repository

import (
	"tt/testtask"

	"github.com/jmoiron/sqlx"
)

type User interface {
	GetByPas(serie, number int) (testtask.Users, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserDB(db),
	}
}
