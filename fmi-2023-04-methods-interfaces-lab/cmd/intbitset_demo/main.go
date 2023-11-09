package main

import (
	"fmt"
	"github.com/iproduct/coursego/fmi-2023-04-methods-interfaces-lab/intbitset"
)

func main() {
	set := intbitset.New()
	fmt.Println(set.String())
	set.Add(42)
	set.Add(129)
	set.Add(95)
	fmt.Println(set.String())
	fmt.Println(set.BitString())
	fmt.Println(set.Has(129))
	fmt.Println(set.Has(135))
	fmt.Println(set.Has(1240))
}
