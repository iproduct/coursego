package main

import "fmt"

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
		case x := <-ch1:
			fmt.Println("Channel 1 sent", x)
		case y := <-ch2:
			fmt.Println("Channel 2 sent", y)
		}
	}
	fmt.Println("Program finished normally.")
}
