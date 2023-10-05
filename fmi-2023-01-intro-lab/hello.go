package main

import (
	"fmi-2023-01-intro-lab/stringutil"
	"fmt"
	"rsc.io/quote"
)

func main() {
	s := "Hello Go World"
	fmt.Println(s)
	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(goquote))
}
