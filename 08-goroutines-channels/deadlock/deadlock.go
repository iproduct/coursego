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
	<-chB
	chA <- 1
	fmt.Println("Demo complete")
}

