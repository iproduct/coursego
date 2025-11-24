package main

import "fmt"

func StringProducer(msg string, n int) <-chan string {
	ch := make(chan string)
	go func(msg string, n int) {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
		}
	}(msg, n)
	return ch
}

func main() {
	p1 := StringProducer("hi", 20)
	for ok, msg := true, ""; ok; {
		msg, ok = <-p1
		fmt.Println(msg)
	}
}
