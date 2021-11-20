package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f *F) M() {
	fmt.Println(*f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	x := F(math.Pi)
	i = &x
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%#v, %T)\n", i, i)
}
