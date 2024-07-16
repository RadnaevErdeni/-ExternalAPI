package repository

import (
	"fmt"
	"tt/testtask"

	"github.com/jmoiron/sqlx"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{db: db}
}

func (r *UserDB) GetByPas(serie, number int) (testtask.Users, error) {
	var user testtask.Users
	query := fmt.Sprintf(`SELECT us.passport_serie,us.passport_number,us.surname,us.name,us.patronymic,us.address FROM %s us WHERE us.passport_serie = $1 AND us.passport_number = $2`,
		usersTable)
	err := r.db.Get(&user, query, serie, number)
	return user, err
}
