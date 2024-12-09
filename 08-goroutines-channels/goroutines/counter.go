package main

import (
	"fmt"
	"time"
)

var count = 0

func counter() {
	for i := 0; i < 1000; i++ {
		go func() {
			count++ // Wrong!
		}()
	}
}

func main() {
	counter()
	time.Sleep(time.Second)
	fmt.Println(count)
}
