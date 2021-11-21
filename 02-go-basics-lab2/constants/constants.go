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
	//userRoles := 12
	var s Stringer
	s = userRoles
	fmt.Printf("User Roles: %s, \n%#[1]T -> %#[1]v\n", s)
	x := make([]int, 1, 5)
	fmt.Printf("%p -> %#v, len=%d, cap=%d\n", &x[0], x, len(x), cap(x))
	x = append(x, 42, 108)
	fmt.Printf("%p -> %#v, len=%d, cap=%d\n", &x[0], x, len(x), cap(x))
	var xp []int
	fmt.Printf("%p -> %#v, len=%d, cap=%d\n", &xp, xp, len(xp), cap(xp))
	xp = append(xp, 42, 108)
	xp = append(xp, 512)
	fmt.Printf("%p -> %#v, len=%d, cap=%d\n", &(xp)[0], xp, len(xp), cap(xp))

	a := [5]int32{'h', 'e', 'l', 'l', '🔦'}
	fmt.Printf("%p -> %#U, len=%d, cap=%d\n", &a[0], a, len(a), cap(a))
	sl := a[:]
	fmt.Printf("%p -> %#U, len=%d, cap=%d\n", &sl[0], sl, len(sl), cap(sl))
	mystr := string(sl)
	fmt.Printf("%s, len=%d \n", mystr, len(mystr))

}
