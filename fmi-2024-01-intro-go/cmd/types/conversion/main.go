package main

import "fmt"

type MyString struct {
	Data []rune
}

func (mystring MyString) String() string {
	return string(mystring.Data)
}

func (mystring *MyString) Append(value string) {
	(*mystring).Data = append((*mystring).Data, []rune(value)...)
}

func (mystring *MyString) Prepend(value string) {
	(*mystring).Data = append([]rune(value), (*mystring).Data...)
}

func (mystring *MyString) Len() int {
	return len((*mystring).Data)
}

func main() {
	var ms MyString
	for i := 0; i < 26; i++ {
		ms.Prepend(string(i + 'A'))
	}
	fmt.Printf("%s\n", ms)
	fmt.Println(ms.Len())

}
