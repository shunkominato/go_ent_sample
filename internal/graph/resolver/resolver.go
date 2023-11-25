package resolver

import (
	"go-gql-sample/app/ent"
	"go-gql-sample/app/internal/graph"
	"go-gql-sample/app/internal/graph/model"

	"github.com/99designs/gqlgen/graphql"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	todos []*model.Todo
	client *ent.Client
}

func NewResolver(client *ent.Client) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
			Resolvers: &Resolver{
			client: client,
		},
	})
}