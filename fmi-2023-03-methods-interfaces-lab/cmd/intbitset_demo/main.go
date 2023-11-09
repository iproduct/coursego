package main

import (
	"fmi-2023-03-methods-interfaces-lab/intbitset"
	"fmt"
)

func main() {
	set := intbitset.New()
	fmt.Println(set.String())
	set.Add(42)
	set.Add(129)
	set.Add(95)
	fmt.Println(set.String())
	fmt.Println(set.BitString())
}
