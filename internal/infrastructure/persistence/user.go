package persistence

import (
	"context"
	"go-gql-sample/app/ent"
	"go-gql-sample/app/ent/user"
	domain "go-gql-sample/app/internal/dmain/entity/user"
	"log"
)

type UserRepository struct{
	client *ent.Client
}

var _ domain.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (ur *UserRepository) Get(ctx context.Context, userIds []int) ([]domain.UserReadModel, error) {
	log.Printf("------------rep 111-----------")
	userList ,err := ur.client.User.Query().Where(user.IDIn(userIds...)).All(ctx)
	if err != nil {
		return []domain.UserReadModel{}, err
	}
	
	log.Printf("%v", userList)

	var result []domain.UserReadModel
	for _, u := range userList {
		result = append(result, *domain.NewUserReadModel(u.ID,u.Age,u.Name))
	}

	return result, nil;
}