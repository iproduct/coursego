package main

import (
	"fmt"
)

type mutable []int

func main() {
	c := make(chan mutable)

	go func() {
		m1 := mutable{42}
		m1[0] = 342
		fmt.Printf("In producer 1: %#v: %p\n", m1, &m1)
		c <- m1
		m1 = mutable{242}
		c <- m1
		//time.Sleep(1 * time.Second)
		//fmt.Printf("In producer 2: %#v: %p\n", m1, &m1)
		//m1[0] = 108 // Don't do this
		//c <- m1p
		close(c)
	}()
	for m := range c {
		fmt.Printf("%#v\n", m)
		//m[0] = 512
		fmt.Printf("In main: %#v: %p\n", m, &m)
	}
}
