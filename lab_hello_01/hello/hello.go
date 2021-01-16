package main

import (
	"fmt"
	"lab_hello_01/stringutil"
	"rsc.io/quote"
	"time"
)

func main() {
	s := "Hello Go World"
	fmt.Println(s, time.Now())
	quote := quote.Go()
	fmt.Println(quote)
	fmt.Println(stringutil.Reverse(quote))
}
