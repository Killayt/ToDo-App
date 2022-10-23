package service

import "github.com/Killayt/ToDo-App/pkg/repos"

type Auth struct {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Auth
	TodoList
	TodoItem
}

func NewService(repos *repos.Repo) *Service {
	return &Service{}
}
