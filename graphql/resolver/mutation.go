package resolver

import (
	"context"
	"go-clean-architecture/model"
	"go-clean-architecture/service"
)

type mutation struct{ service *service.Service }

func (m mutation) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	return m.service.Todo.Create(ctx, input)
}
