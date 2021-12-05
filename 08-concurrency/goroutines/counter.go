package main

import (
	"fmt"
	"time"
)

func counter() {
	count := 0

	for i := 0; i < 1000; i++ {
		go func() {
			count++ // Wrong!
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(count)
}
