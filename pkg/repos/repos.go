package repos

type Auth struct {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repo struct {
	Auth
	TodoList
	TodoItem
}

func NewService() *Repo {
	return &Repo{}
}
