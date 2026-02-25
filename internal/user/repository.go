package user

import "github.com/jmoiron/sqlx"

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}
