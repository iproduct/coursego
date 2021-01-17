package main

import "fmt"

type Role int

const (
	User Role = 1 << iota
	Manager
	Admin
	RoleMask = (1 << (iota)) - 1
)

func (r Role) String() string {
	switch r {
	case User:
		return "User"
	case Manager:
		return "Manager"
	case Admin:
		return "Admin"
	default:
		return "Invalid role"
	}
}

func main() {
	fmt.Printf("%s : %[1]V, Mask: %b", Admin, RoleMask)
}

// fmt package defines interface Stringer as:
//type Stringer interface {
//	String() string
//}
