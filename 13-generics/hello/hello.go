package main

import (
	"fmt"
	"log"
)

func Print[T any](thing T) {
	fmt.Println(thing)
}

//type Stringer interface {
//	String() string
//}
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

func Join[E any](things []E, separator string) (result string, err error) {
	result = ""
	lastIndex := len(things) - 1
	for i, v := range things {
		str, err := Stringify(v)
		if err != nil {
			return "", err
		}
		result += str
		if i < lastIndex {
			result += separator
		}
	}
	return result, nil
}

func main() {
	Print("Hello!")
	Print(42)
	Print(true)
	output, err := Join([]string{"a", "b", "c"}, ", ")
	if err != nil {
		log.Fatalln(err)
	}
	Print(output)

	users := []User{
		{"1", "john", "john123"},
		{"2", " mary", "mary123"},
		{"3", "hristo", "hristo123"},
	}
	output, err = Join(users, "\n")
	if err != nil {
		log.Fatalln(err)
	}
	Print(output)

	output, err = Join([]int{1, 2, 3}, ", ")
	if err != nil {
		log.Fatalln(err)
	}
	Print(output)

}
