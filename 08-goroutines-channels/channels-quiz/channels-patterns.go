package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() { ch <- "hi" }()
	select {
	case case1 := <-ch:
		fmt.Printf("case1: %s\n", case1)
	case case2 := <-ch:
		fmt.Printf("case2: %s\n", case2)
	case <-time.After(time.Nanosecond):
		fmt.Printf("Timeout after 1 ns\n")
	default:
		fmt.Printf("No activity ...\n")
	}
}
