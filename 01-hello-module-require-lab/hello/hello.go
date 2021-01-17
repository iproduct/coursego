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
	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(goquote))
}
