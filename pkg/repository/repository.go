package repository

import (
	"github.com/jmoiron/sqlx"
	"restAPI"
)

type Authorization interface {
	CreateUser(user restAPI.User) (int, error)
	GetUser(username, password string) (restAPI.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
