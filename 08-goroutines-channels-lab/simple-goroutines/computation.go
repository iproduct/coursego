package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan struct{})
	go compute("Long computation", done)
	//time.Sleep(400 * time.Second)
	v, ok := <-done
	fmt.Println(v, ok)
	fmt.Println("Demo complete.")
}

func compute(msg string, done chan<- struct{}) {
	defer close(done)
	for i := 0; i < 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
