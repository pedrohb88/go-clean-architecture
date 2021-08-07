package repository

import (
	"context"
	"fmt"
	"go-clean-architecture/model"
)

type todoRepository struct{}

type Todo interface {
	All(ctx context.Context) ([]*model.Todo, error)
	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
}

func NewTodo() *todoRepository {
	return &todoRepository{}
}

func (r *todoRepository) All(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println("Fetching all todos from database")
	return []*model.Todo{
		{
			ID:   "1",
			Text: "First todo",
			Done: false,
		},
		{
			ID:   "2",
			Text: "Second todo",
			Done: true,
		},
	}, nil
}

func (r *todoRepository) Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	fmt.Println("Creating new todo on database")
	return &model.Todo{
		ID:   "novo id",
		Text: input.Text,
		Done: false,
	}, nil
}
