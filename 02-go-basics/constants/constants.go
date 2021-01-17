package main

import "fmt"

type Role  int

const (
	User Role = 1 << iota
	Manager
	Admin
	Customer
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
	case Customer:
		return "Customer"
	default:
		return "Invalid role"
	}
}

func main() {
	fmt.Printf("%s - %[1]d : %[1]V, Mask: %b\n", User, RoleMask)
	fmt.Printf("%s - %[1]d: %[1]V, Mask: %b\n", Manager, RoleMask)
	fmt.Printf("%s - %[1]d : %[1]V, Mask: %b\n", Admin, RoleMask)
	fmt.Printf("%s - %[1]d : %[1]V, Mask: %b\n", Customer, RoleMask)
}

// fmt package defines interface Stringer as:
type Stringer interface {
	String() string
}
