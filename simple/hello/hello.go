package main

import "fmt"
import "github.com/iproduct/coursego/simple/stringutil"
import "rsc.io/quote"

func main() {
	s := "Hello Go World!"
	fmt.Println(s)
	fmt.Println(stringutil.Reverse(s))
	s = quote.Go()
	fmt.Println(s)
	fmt.Println(stringutil.Reverse(s))
}
