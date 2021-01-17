package main

import (
	"fmt"
	"github.com/iproduct/coursego/01-hello-lab/stringutil"
	"rsc.io/quote"
	"time"
)

func main() {
	s := "abc日本語"
	for index, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	fmt.Println(s, time.Now())
	quote := quote.Go()
	fmt.Println(quote)
	fmt.Println(stringutil.Reverse(s))
}
