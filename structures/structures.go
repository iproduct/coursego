package main

import (
	"fmt"
	"github.com/iproduct/course-go/enums"
	"strings"
)

// Name type
type Name string

// ConvertionCase type
type ConvertionCase int

// User struct models a registered user of the service
type User struct {
	id       int
	Name     Name
	Username Name
	Password string
	Roles    enums.Role
	Status   enums.Status
}

// Name case for conversion and printing
const (
	UpperCase = iota
	LowerCase = iota
	TitleCase = iota
)

func (name Name) convert(convCase ConvertionCase) Name {
	var result string
	switch convCase {
	case UpperCase:
		result = strings.ToUpper(string(name))
	case LowerCase:
		result = strings.ToLower(string(name))
	case TitleCase:
		result = strings.ToTitle(string(name))
	default:
		result = string(name)
	}
	return Name(result)
}

func (u *User) String(convCase ConvertionCase) string {
	return fmt.Sprintf("Name: %s, Username: %s, in Role: %s, Status: %s",
		u.Name.convert(convCase), u.Username, u.Roles.String(), u.Status.String())
}

func main() {
	user := User{1, "John Smith", "john", "john123", enums.Admin, enums.Active}
	fmt.Printf("%s", user.String(UpperCase))
}
