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

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i) //(&{Hello}, *intbitset_main.T)
	i.M()       //Hello

	i = F(math.Pi)
	describe(i) //(3.141592653589793, intbitset_main.F)
	i.M()       //3.141592653589793
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
