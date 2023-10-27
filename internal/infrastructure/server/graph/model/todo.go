package model

type Todo struct {
	ID           string `json:"id"`
	Todo         string `json:"todo"`
	TodoStatusID int    `json:"todoStatusId"`
	User         *User  `json:"user"`
}
