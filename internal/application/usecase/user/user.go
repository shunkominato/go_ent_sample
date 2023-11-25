package user

import (
	"context"
	"go-gql-sample/app/internal/dmain/entity/user"
)

type IUserUsecase interface {
	Get(ctx context.Context, userIds []int) ([]user.UserReadModel, error)
}

type userUsecase struct {
	ur user.IUserRepository
}

func NewUserUsecase(ur user.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Get(ctx context.Context, userIds []int) ([]user.UserReadModel, error) {
	userList , err := uu.ur.Get(ctx, userIds)
	if err != nil {
		return userList, err
	}
	return userList, nil;
}