package main

import "fmt"

type Status int

const (
	Registered Status = iota + 1
	Active
	Disabled
)

type Stringer interface {
	String() string
}

func (s Status) String() string { // Duck typing
	switch s {
	case Registered:
		return "Registered"
	case Active:
		return "Active"
	case Disabled:
		return "Disabled"
	default:
		return "invalid status"
	}
}

func main() {
	userStatus := Disabled
	fmt.Printf("User Staus: %#v -> %[1]s", userStatus)
}
