package main

import "fmt"

// Role type
type Role int

// Roles enum
const (
	User Role = 1 << iota
	Manager
	Admin
	RoleMask = (1 << (iota)) - 1
)

// Returns string representation of the Role
func (r Role) String() string {
	result := ""
	if r&User == User {
		result += "User "
	}
	if r&Manager == Manager {
		result += "Manager "
	}
	if r&Admin == Admin {
		result += "Admin"
	}
	return result
}

// Status type
type Status int

// User statuses enum
const (
	Registered Status = iota
	Active
	Disabled
)

// Returns string representation of the Role
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

func main() {
	fmt.Printf("Roles[%T]: %[1]v, %v, %v\n", User+Manager+Admin, Manager, Admin)
	fmt.Printf("RoleMask: %b \n", RoleMask)

	fmt.Printf("Statuses: %v, %v, %v\n", Registered, Active, Disabled)

}
