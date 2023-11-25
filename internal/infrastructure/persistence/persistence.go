package persistence

import (
	"go-gql-sample/app/ent"
	"go-gql-sample/app/internal/dmain/entity/todo"
)

type Repositories struct {
	client *ent.Client
	Todo todo.ITodoRepository
}

func NewRepositories(client *ent.Client) (*Repositories, error) {
	return &Repositories{
		client: client,
		Todo: NewTodoRepository(client),
	}, nil
}
