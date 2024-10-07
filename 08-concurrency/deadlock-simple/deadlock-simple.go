package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1
	ch <- 2
	close(ch)

	//fmt.Println(<-ch)
	for v := range ch {
		fmt.Println(v)
	}
}
