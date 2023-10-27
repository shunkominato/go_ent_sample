package repository

import (
	"go-gql-sample/app/ent"
)

type Repository struct{
	client *ent.Client
}