package graph

//go:generate go run github.com/99designs/gqlgen generate
import (
	"go-gql-sample/app/ent"
	"go-gql-sample/app/internal/infrastructure/server/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
	todoStatus []*model.TodoStatus
	client *ent.Client
}