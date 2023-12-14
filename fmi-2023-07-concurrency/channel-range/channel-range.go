package main

import "fmt"

func make_ch() <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Sending message number %d", i)
		}
	}()
	return ch
}

func main() {
	ch := make_ch()
	//ch <- "aaaa"
	var val string
	for val = range ch {
		fmt.Printf("Receiving: %#v\n", val)
	}
	print("Demo finished.")
	//ch <- fmt.Sprintf("... but I have a question ...") // results in panic}
}
