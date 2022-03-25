package service

import (
	"restAPI"
	"restAPI/pkg/repository"
)

type Authorization interface {
	CreateUser(user restAPI.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list restAPI.TodoList) (int, error)
	GetAll(userId int) ([]restAPI.TodoList, error)
	GetById(userId, listId int) (restAPI.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
