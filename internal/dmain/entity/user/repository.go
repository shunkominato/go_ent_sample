package user

import (
	"context"
)

type IUserRepository interface {
	Get(ctx context.Context, userIds []int) ([]UserReadModel, error)
}