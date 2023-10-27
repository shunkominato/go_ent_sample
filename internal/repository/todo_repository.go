package repository

import (
	"context"
	"errors"
	"go-gql-sample/app/ent"
)
type TodoRepository struct {
	Client *ent.Client
}

func NewTodoRepository(client *ent.Client) *TodoRepository {
	return &TodoRepository{Client: client}
}

func (repo *TodoRepository) GetTodoList(ctx context.Context) ([]*ent.Todo, error) {
	todoList, err := repo.Client.Todo.Query().All(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return todoList, err
}