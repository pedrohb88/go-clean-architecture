package service

import (
	"context"
	"go-clean-architecture/model"
	"go-clean-architecture/repository"
)

type TodoInterface interface {
	GetAll(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
}

type todo struct {
	todoRepository repository.Todo
}

func (t todo) GetAll(ctx context.Context) ([]*model.Todo, error) {
	return t.todoRepository.All(ctx)
}

func (t todo) Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	return t.todoRepository.Create(ctx, input)
}
