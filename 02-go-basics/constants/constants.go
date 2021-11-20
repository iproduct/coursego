package main

import "fmt"

type Role byte

const (
	User Role = 1 << iota
	Manager
	Admin
	Customer
	RoleMask = (1 << (iota)) - 1
)

type Status int

const (
	Registered Status = iota
	Active
	Disabled
)

// Returns string representation of the Status
func (r Status) String() string {
	switch r {
	case Registered:
		return "Registered"
	case Active:
		return "Active"
	case Disabled:
		return "Disabled"
	default:
		return "Invalid status"
	}
}

func (r Role) String() string {
	result := ""
	if User&r > 0 {
		result += "User "
	}
	if Manager&r > 0 {
		result += "Manager "
	}
	if Admin&r > 0 {
		result += "Admin "
	}
	if Customer&r > 0 {
		result += "Customer"
	}
	if r&RoleMask == 0 {
		result = "Invalid role"
	}
	return result
}

func main() {
	fmt.Printf("%s - %[1]d : %[1]T, Mask: %#b\n", User+Manager, RoleMask)
	fmt.Printf("%s - %[1]d: %[1]T, Mask: %b\n", Manager+Admin, RoleMask)
	fmt.Printf("%s - %[1]d : %[1]T, Mask: %b\n", Admin, RoleMask)
	fmt.Printf("%s - %[1]d : %[1]T, Mask: %b\n\n", Customer, RoleMask)

	fmt.Printf("%s - %[1]d : %[1]T\n", Registered)
	fmt.Printf("%s - %[1]d: %[1]T\n", Active)
	fmt.Printf("%s - %[1]d : %[1]T\n", Disabled)

}

// fmt package defines interface Stringer as:
type Stringer interface {
	String() string
}
