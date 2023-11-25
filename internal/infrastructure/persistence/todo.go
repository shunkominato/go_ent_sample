package persistence

import (
	"context"
	"go-gql-sample/app/ent"
	"go-gql-sample/app/ent/todo"
	domain "go-gql-sample/app/internal/dmain/entity/todo"
	"go-gql-sample/app/internal/dmain/entity/user"
	"log"
)

type TodoRepository struct {
	client *ent.Client
}

var _ domain.ITodoRepository = (*TodoRepository)(nil)

func NewTodoRepository(client *ent.Client) *TodoRepository {
	return &TodoRepository{
		client: client,
	}
}

func (tr *TodoRepository) Create(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	createdTodo, err := tr.client.Todo.
		Create().
		SetName(todo.Name).
		SetText(todo.Text).
		SetDone(todo.Done).
		SetOwnerID(1).
		Save(ctx)
		log.Printf("%v", createdTodo)

	if err != nil {
		log.Printf("%s",err)
		return todo, err
	}

	domainTodo := &domain.Todo{
		ID:        createdTodo.ID,
		Name:      createdTodo.Name, 
		Text:      createdTodo.Text,
		Done:      createdTodo.Done,
		UserID:    createdTodo.UserID,
	}


	return *domainTodo, nil
}

func (tr *TodoRepository) Get(ctx context.Context, todoIds []int) ([]domain.TodoReadModel, error) {
	todoList, err := tr.client.Todo.Query().Where(todo.IDIn(todoIds...)).WithOwner().All(ctx)
	if err != nil {
		log.Printf("%s",err)
		return []domain.TodoReadModel{}, err
	}
	log.Printf("------------rep 111-----------")
	log.Printf("%v", todoList)

	var result []domain.TodoReadModel
	for _, t := range todoList {
		log.Printf("%v", t.Edges.Owner)
		u := t.Edges.Owner
		result = append(result, domain.TodoReadModel{
			ID:   t.ID,
			Name: t.Name,
			Text: t.Text,
			Done: t.Done,
			User: *user.NewUserReadModel(u.ID,u.Age,u.Name),
		})
	}

	return result, nil
}