package main

import (
	"fmt"
	"time"
)

func main() {
	compute("compute1!")
	//compute("compute2!")
}

func compute(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(1 * time.Second)
	}
}
