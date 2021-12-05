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

	ok1, ok2 := true, true
	for {
		select {
		case x, ok := <-ch1:
			ok1 = ok
			if ok {
				fmt.Println("Channel 1 sent", x)
			}
		case y, ok:= <-ch2:
			ok2 = ok
			if ok {
				fmt.Println("Channel 2 sent", y)
			}
		}
		if !(ok1 || ok2) {
			break
		}
	}
	fmt.Println("Program finished normally.")
}
