package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	for v := range ch {
		fmt.Println(v)
	}
}
