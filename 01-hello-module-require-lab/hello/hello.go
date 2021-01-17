package main

import (
	"fmt"
	"github.com/iproduct/coursego/01-hello-lab/stringutil"
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
