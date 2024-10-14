package main

import "fmt"

type MyString []rune

func (mystring *MyString) String() string {
	return string(*mystring)
}

func (mystring *MyString) Append(value string) {
	*mystring = append(*mystring, []rune(value)...)
}

func (mystring *MyString) Prepend(value string) {
	*mystring = append([]rune(value), *mystring...)
}

func (mystring *MyString) Len() int {
	return len(*mystring)
}

func main() {
	var ms MyString
	for i := 0; i < 26; i++ {
		ms.Prepend(string(i + 'A'))
	}
	fmt.Println(string(ms))
	fmt.Println(ms.Len())

}
