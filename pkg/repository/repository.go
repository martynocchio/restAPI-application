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
	Create(userId int, list restAPI.TodoList) (int, error)
	GetAll(userId int) ([]restAPI.TodoList, error)
	GetById(userId, listId int) (restAPI.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input restAPI.UpdateListInput) error
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
		TodoList:      NewTodoListPostgres(db),
	}
}
