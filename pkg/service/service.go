package service

import "github.com/Killayt/ToDo-App/pkg/repository"

type Auth interface{}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Auth
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
