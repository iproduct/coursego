package model

// User models a REST API user
type User struct {
	ID   int    `json:"id" validate:"numeric,gte=0"`
	Name string `json:"name" validate:"required,min=5,max=30"`
	Age  int    `json:"age" validate:"required,numeric,gte=0,lte=130"`
}
