package user

import (
	"database/sql"
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) User {
	return User{db}
}

func (u User) Create() {

}
