package main

import (
	"fmt"
	"rsc.io/quote"
	"github.com/iproduct/coursegopro/01-intro-lab/stringutil"
)

func main() {
	s := "abc日本語"
	r := []rune(s)
	for i:=0; i < len(r); i++ {
		fmt.Printf("%#U starts at byte position %d\n", r[i], i)
	}

	for index, runeVal := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeVal, index)
	}
	
	fmt.Println(s)
	var goquote string = quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(goquote))	
}
