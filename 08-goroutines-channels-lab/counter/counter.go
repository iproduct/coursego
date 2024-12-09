package main

import "fmt"

func main() {
	counter := 0
	for i := 0; i < 10000; i ++ {
		go func() {
			counter ++
		}()
	}
	fmt.Printf("Counter = %d\n", counter)
}
