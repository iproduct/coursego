package main

import "fmt"

type Role int

const (
	User Role = 1 << iota
	Manager
	Admin
	Customer
	RoleMask = (1 << iota) - iota
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

type Status int

const (
	Registered = iota
	Active
	Disabled
)

func main() {
	fmt.Printf("%s : %[1]T\n", User)
	fmt.Printf("%s : %[1]T\n", Manager)
	fmt.Printf("%s : %[1]T\n", Admin)
	fmt.Printf("%s : %[1]T\n", Customer)
}
