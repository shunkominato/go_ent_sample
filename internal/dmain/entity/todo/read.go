package todo

import "go-gql-sample/app/internal/dmain/entity/user"


type TodoReadModel struct {
	ID int 
	Name string `validate:"required,lte=5"`
	Text string `validate:"required,lte=5"`
	Done bool
	User user.UserReadModel
}

func NewTodoReadModel(id int, name string, text string, done bool, user user.UserReadModel) (*TodoReadModel){
	return &TodoReadModel{
		ID: id,
		Name: name,
		Text: text,
		Done: done,
		User: user,
	}
}