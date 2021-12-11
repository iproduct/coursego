package model

type User struct {
	ID       int    `json:"id" validate:"numeric,gte=0"`
	Name     string `json:"name" validate:"required,min=5,max=30"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Age      int    `json:"age" validate:"required,numeric,gte=0,lte=130"`
	Active   bool   `json:"active"`
}
