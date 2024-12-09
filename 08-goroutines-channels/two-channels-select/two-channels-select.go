package main

import (
	"fmt"
	"time"
)

func sendTo(c chan<- int, iter int) {
	for i := 0; i <= iter; i++ {
		c <- i
		//time.Sleep(time.Duration(100 * time.Millisecond))
		<-time.After(5000 * time.Millisecond)
	}
	close(c)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendTo(ch1, 5)
	go sendTo(ch2, 10)

	for {
		select {
		case x := <-ch1:
			fmt.Println("Channel 1 sent", x)
		case y := <-ch2:
			fmt.Println("Channel 2 sent", y)
			//default:
			//	fmt.Println("Nothing received")
			//	time.Sleep(time.Duration(500 * time.Millisecond))
		}
	}
	fmt.Println("Program finished normally.")
}
