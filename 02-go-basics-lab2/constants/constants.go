package main

import "fmt"

type Roles int

const (
	User Roles = 1 << iota
	Customer
	Manager
	Admin
	RoleMask = (1 << iota) - 1
)

type Stringer interface {
	String() string
}

func (r Roles) String() string { // Duck typing
	r = r & RoleMask
	result := ""
	if r&User > 0 {
		result = result + "User "
	}
	if r&Customer > 0 {
		result = result + "Customer "
	}
	if r&Manager > 0 {
		result = result + "Manager "
	}
	if r&Admin > 0 {
		result = result + "Admin "
	}
	return result
}

func main() {
	userRoles := Customer + Manager + Admin
	fmt.Printf("User Roles: %s", userRoles)
}
