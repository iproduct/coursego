package main

import (
	"fmt"
)

func goroutine2(chA, chB chan int) {
	<-chA
	chB <- 1
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	go goroutine2(chA, chB)
	<-chB
	chA <- 1
	fmt.Println("Demo complete")
}
