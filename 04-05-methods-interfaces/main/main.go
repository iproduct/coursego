package main

import (
	"fmt"
	"github.com/iproduct/coursego/04-05-methods-interfaces/intbitset"
)

func main() {
	var x, y intbitset.IntBitSet

	x.Add(2)
	x.Add(42)
	x.Add(120)
	y.Add(9)
	y.Add(42)
	y.Add(145)

	fmt.Println("42 in x:", x.Has(42))
	fmt.Println("42 in y:", y.Has(143))
	fmt.Printf("x = %s\n", x.String())
	fmt.Printf("y = %s\n", y.String())
}
