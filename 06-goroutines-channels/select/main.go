package main

import "fmt"

func main() {
	ch := make(chan string)
	select {
	case <-ch:
		fmt.Println("Data received")
	}
	fmt.Println("Demo finished")
	ch <- "hi"
}
