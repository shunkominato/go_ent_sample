package todo

import (
	"go-gql-sample/app/internal/graph/validation"
	"log"
)

type Todo struct {
	ID int 
	Name string `validate:"required,lte=5"`
	Text string `validate:"required,lte=5"`
	Done bool
	UserID int
}

type SaveTodoInputDto struct {
	ID int
	Text string
	UserID string  
}

type GetTodoInputDto struct {
	ID int
}


func NewTodo(input SaveTodoInputDto) (Todo, error) {
	log.Printf("------------ddd-----------")
	todo := Todo{
		ID: input.ID,
		Name: input.Text,
		Text: input.Text,
		Done: true,
		UserID: '1',
	}
	log.Printf("------------eee-----------")

	if err := validation.ValidateInputModel(todo); err != nil {
		return todo, err
	}
	log.Printf("------------fff-----------")

	return todo, nil

}