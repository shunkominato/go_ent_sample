package todo

import (
	"context"
	"log"

	"go-gql-sample/app/internal/dmain/entity/todo"
)

type ITodoUsecase interface {
	Create(ctx context.Context, todoInput *todo.SaveTodoInputDto) (todo.Todo, error)
	Get(ctx context.Context, todoInput []int)  ([]todo.TodoReadModel, error)
}

type todoUsecase struct {
	tr todo.ITodoRepository
}

func NewTodoUsecase(tr todo.ITodoRepository) ITodoUsecase {
	return &todoUsecase{tr}
}

func (tu *todoUsecase) Create(ctx context.Context, todoInput *todo.SaveTodoInputDto)(todo.Todo, error) {
	log.Printf("------------usecase 111-----------")
	todoEntity, err := todo.NewTodo(*todoInput)
	if err != nil {
		return todoEntity, err
	}
	log.Printf("------------usecase 222-----------")
	createdTodo, err := tu.tr.Create(ctx, todoEntity)
	if err != nil {
		return createdTodo, err
	}
	log.Printf("------------usecase 333-----------")
	return createdTodo, nil

}

func (tu *todoUsecase) Get(ctx context.Context, todoIds []int) ([]todo.TodoReadModel, error) {
	log.Printf("------------usecase 111-----------")
	todoList, err := tu.tr.Get(ctx, todoIds)
	if err != nil {
		return todoList, err
	}
	return todoList, nil;
}