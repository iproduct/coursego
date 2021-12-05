package main

import (
	"fmt"
)

func sendTo(c chan<- int, iter int) {
	for i := 0; i <= iter; i++ {
		c <- i
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
		case x, ok := <-ch1:
			if ok {
				fmt.Println("Channel 1 sent", x)
			} else {
				ch1 = nil
			}
		case y, ok := <-ch2:
			if ok {
				fmt.Println("Channel 2 sent", y)
			} else {
				ch2 = nil
			}
			//default:
			//	fmt.Println("Not active ...",)
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("Program finished normally.")
}
