package main

import "fmt"

func fn(mm *(map[int]int)) {
	*mm = make(map[int]int)
}

func main() {
	var mm map[int]int
	fmt.Printf("%#v\n", mm)
	fn(&mm)
	fmt.Println(mm == nil)
	fmt.Printf("%#v\n", mm)

}
