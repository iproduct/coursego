package main

import (
	"fmt"
)

func myfunc(chA, chB chan int) {
	<-chA
	chB <- 1
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	go myfunc(chA, chB)
	chA <- 1
	<-chB
	fmt.Println("Demo complete")
}
