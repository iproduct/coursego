package main

import (
	"fmt"
	"github.com/iproduct/coursego/03-types-lab/enums"
	"regexp"
	"strings"
)

// ConvertionCase type
type ConvertionCase int

// Name case for conversion and printing
const (
	UpperCase = iota
	LowerCase
	TitleCase
)

// Name type
type Name string

func (name Name) convert(convCase ConvertionCase) Name {
	var result string
	switch convCase {
	case UpperCase:
		result = strings.ToUpper(string(name))
	case LowerCase:
		result = strings.ToLower(string(name))
	case TitleCase:
		{
			r := regexp.MustCompile("\\s+")
			names := r.Split(string(name), -1)
			// names := strings.Split(string(name), " ")
			for pos, name := range names {
				result += strings.Title(name)
				if pos < len(names)-1 {
					result += " | "
				}
			}
		}
		// result = strings.Title(string(name))
	default:
		result = string(name)
	}
	return Name(result)
}

// User struct models a registered user of the service
type User struct {
	id       int
	Name     Name
	Username Name
	Password string
	Roles    enums.Role
	Status   enums.Status
}

func (u *User) String(convCase ConvertionCase) string {
	return fmt.Sprintf("Name: %s, Username: %s, in Role: %s, Status: %s",
		u.Name.convert(convCase), u.Username, u.Roles.String(), u.Status.String())
}

// func (r enums.Role) Valid() bool{
// 	return r & enums.RoleMask != 0;
// }

func main() {
	user := User{1, "john 	 smith", "john", "john123", enums.Admin, enums.Active}
	fmt.Printf("%s", user.String(TitleCase))
}
