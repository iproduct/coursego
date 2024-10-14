package main

import "fmt"

// Status type
type Status int

// User statuses enum
const (
	Registered Status = iota + 1
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
	status := Registered
	fmt.Println(byte(status + 1))
}
