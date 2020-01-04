package main

import (
	"time"
	"strconv"
	"fmt"
)

var global int

func main() {
	global = 42
	// var local int = 5
	local, n := 5, 12
	fmt.Printf("global = %d\nlocal = %d\n", global, local+n)
	var cp1 complex128 = 12 + 3i
	cp2 := 5 + 7i
	fmt.Printf("%#v\n", cp1 + cp2)
	str := "Hello"
	fmt.Printf("%-20.20s -> %5d\n", str, n)

	ch := make(chan interface{})
	go func() {
		defer close(ch)
		var i int64
		for i = 0;  i < 10; i++ {
			ch <- "message_" + strconv.FormatInt(i, 10);
			time.Sleep(time.Second)
		}
	}()

	for {
		result, ok := <-ch
		if ok {
			fmt.Printf("%v\n", result)
		} else {
			fmt.Printf("END\n")
			break
		}
	}
}
