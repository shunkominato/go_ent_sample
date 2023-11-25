package todo

import (
	"context"
)

type ITodoRepository interface {
	Create(ctx context.Context, todo Todo) (Todo, error)
	Get(ctx context.Context, todoIds []int) ([]TodoReadModel, error)
}