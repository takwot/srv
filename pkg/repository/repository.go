package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/takwot/srv"
)

const (
	usersTable = "users"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type Film interface {
}

type Films interface {
}

type Repository struct {
	Authorization
	Film
	Films
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
