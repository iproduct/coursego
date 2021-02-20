package model

var TodoList []Todo

type Todo struct {
	ID   string `json:id`
	Text string `json:text`
	Done bool   `json:done`
}
