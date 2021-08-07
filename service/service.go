package service

import "go-clean-architecture/repository"

type Service struct {
	Todo TodoInterface
}

func New() *Service {

	todoService := &todo{
		todoRepository: repository.NewTodo(),
	}

	return &Service{
		Todo: todoService,
	}
}
