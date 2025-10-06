package main

import (
	"fmt"
	"runtime"
	"time"
)

type mutable []int

func main() {
	runtime.GOMAXPROCS(12)
	c := make(chan mutable)
	go func() {
		m1 := mutable{42}
		c <- m1
		time.Sleep(1 * time.Second)
		m1[0] = 342
		close(c)
	}()
	m1 := make([]int, 1)
	for m := range c {
		time.Sleep(1001 * time.Millisecond)
		copy(m1, m)
		fmt.Printf("%v\n", m1)
	}
}
