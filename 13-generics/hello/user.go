package main

import "fmt"

type User struct {
	ID, Username, Password string
}

func (u User) String() string {
	return fmt.Sprintf("User[ID: %s, Name: %s, Password: %s]", u.ID, u.Username, u.Password)
}
