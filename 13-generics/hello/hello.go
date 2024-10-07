package main

import (
	"fmt"
)

func Print[T any](thing T) {
	fmt.Println(thing)
}

type Stringer interface {
	String() string
}

//
//type StringLike interface {
//	~string
//	String() string
//}

func Stringify(s any) (string, error) {
	switch v := s.(type) {
	case fmt.Stringer:
		return v.String(), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("I can not stringify type %T!\n", v)
	}
}

func Join[S ~[]E, E Stringer](things S, separator string) string {
	result := ""
	lastIndex := len(things) - 1
	for i, v := range things {
		result += v.String()
		if i < lastIndex {
			result += separator
		}
	}
	return result
}

func main() {
	Print("Hello!")
	Print(42)
	Print(true)
	//output := Join([]string{"a", "b", "c"}, ", ")
	//Print(output)

	users := []User{
		{"1", "john", "john123"},
		{"2", " mary", "mary123"},
		{"3", "hristo", "hristo123"},
	}
	Print(Join(users, "\n"))

	//
	//output, err = Join([]int{1, 2, 3}, ", ")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//Print(output)

}
