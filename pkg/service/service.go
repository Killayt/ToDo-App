package service

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

func NewService() *Service {
	return &Service{}
}
