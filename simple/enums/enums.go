package enums

import "fmt"

// Role type
type Role int

// Roles enum
const (
	User     = 1 << iota
	Manager  = 1 << iota
	Admin    = 1 << iota
	RoleMask = (1 << (iota)) - 1
)

// Returns string representation of the Role
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

// Status type
type Status int

// User statuses enum
const (
	Registered = iota
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
	fmt.Printf("Roles: %v, %v, %v\n", User, Manager, Admin)
	fmt.Printf("RoleMask: %b \n", RoleMask)

	fmt.Printf("Statuses: %v, %v, %v\n", Registered, Active, Disabled)

}
