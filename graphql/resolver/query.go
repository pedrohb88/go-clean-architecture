package resolver

import (
	"context"
	"go-clean-architecture/model"
	"go-clean-architecture/service"
)

type query struct{ service *service.Service }

func (q query) Todos(ctx context.Context) ([]*model.Todo, error) {
	return q.service.Todo.GetAll(ctx)
}
