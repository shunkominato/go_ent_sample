package user

type UserReadModel struct
 {
	ID int
	Age int
	Name string `validate:"required,lte=5"`
}

func NewUserReadModel(id int, age int, name string)(*UserReadModel){
	return &UserReadModel{
		ID: id,
		Age: age,
		Name: name,
	}
}

