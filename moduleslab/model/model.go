package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Eamil    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
