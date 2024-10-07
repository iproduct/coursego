package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0

	for i := 0; i < 1000; i++ {
		go func() {
			count++
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(count)
}
