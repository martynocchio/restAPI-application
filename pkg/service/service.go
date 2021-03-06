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
	Delete(userId, listId int) error
	Update(userId, listId int, input restAPI.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item restAPI.TodoItem) (int, error)
	GetAll(userId, listId int) ([]restAPI.TodoItem, error)
	GetById(userId, itemId int) (restAPI.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input restAPI.UpdateItemInput) error
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
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
