package service

import (
	"context"
	"errors"
	"go-gql-sample/app/ent"
	"go-gql-sample/app/ent/todostatus"
	"go-gql-sample/app/internal/infrastructure/server/graph/model"
	"log"
	"strconv"
)

func GetTodoStatus(ctx context.Context, client *ent.Client) (*model.TodoStatus, error) {
	todoStatus, err := client.TodoStatus.Query().Where(todostatus.ID(1)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	log.Print(todoStatus)
	return &model.TodoStatus{
		ID: strconv.Itoa(todoStatus.ID),
		Status: todoStatus.Status,
	}, nil
}