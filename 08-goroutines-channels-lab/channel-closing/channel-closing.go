package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Sending message number %d", i)
		}
		close(ch)
	}()
	var val string
	for ok := true; ok; {
		val, ok = <-ch
		fmt.Printf("Receiving: %#v, %#v\n", val, ok)
	}
	ch <- fmt.Sprintf("... but I have a question ...")
}
