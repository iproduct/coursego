package main

import (
	"fmt"
	"github.com/iproduct/coursego/fmi-2026-01-intro/stringutil"
	"rsc.io/quote"
)

func main() {
	s := "Hello Go World! 你好世界"
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c -> %d\n", s[i], i)
	}
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c -> %d\n", r[i], i)
	}

	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(goquote))
	fmt.Println(stringutil.Reverse(s))

}
