package main

import "fmt"
import "github.com/iproduct/coursego/simple/stringutil"

func main() {
	s := "Hello Go World!"
	fmt.Println(s)
	fmt.Println(stringutil.Reverse(s))
}
