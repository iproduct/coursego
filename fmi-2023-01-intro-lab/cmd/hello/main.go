package main

import (
	"fmt"
	"github.com/iproduct/coursego/fmi-2023-01-intro-lab/stringutil"
	"github.com/iproduct/coursego/fmi-2023-03-methods-interfaces-lab/intbitset"
	"rsc.io/quote"
)

func main() {
	s := "Hello Go World"
	fmt.Println(s)
	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(goquote))
	set := intbitset.New()
}
