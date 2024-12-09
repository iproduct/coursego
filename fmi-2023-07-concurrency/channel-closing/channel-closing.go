package main

import "fmt"

func main() {
	ch := make(chan string, 9)
	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("Sending message number %d", i)
		}
		close(ch)
		close(done)
	}()
	//v, ok := <-done
	//fmt.Printf("%T %v, %t\n", v, v, ok)
	var val string
	for ok := true; ok; {
		val, ok = <-ch
		fmt.Printf("Receiving: %#v, %#v\n", val, ok)
	}
	print("Demo finished.")
	//ch <- fmt.Sprintf("... but I have a question ...") // results in panic}
}
