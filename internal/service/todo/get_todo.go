package todoService

import (
	"context"
	"errors"
	"go-gql-sample/app/ent"
	"go-gql-sample/app/internal/repository"
	slackNotify "go-gql-sample/app/pkg/slack"
)
type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetTodoList(ctx context.Context) ([]*ent.Todo, error) {
	todoList, err := s.repo.GetTodoList(ctx)
	if err != nil {
		slackNotify.ErrorNotificator("", "GetTodoList", err, "")
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return todoList, err
}