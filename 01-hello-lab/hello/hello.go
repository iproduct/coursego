package main

import (
	"fmt"
	"github.com/iproduct/coursego/01-hello-lab/stringutil"
	"github.com/iproduct/coursego/fmi-2023-03-methods-interfaces-lab/intbitset"
	"rsc.io/quote"
	"time"
)

func main() {
	s := "abc日本語"
	for index, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U starts at byte position %d\n", s[i], i)
	}
	fmt.Println(s, time.Now())
	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(s))
	set := intbitset.New()
	fmt.Println(set)
}
