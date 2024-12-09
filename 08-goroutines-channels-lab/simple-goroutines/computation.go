package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go compute("Long computation")
	time.Sleep(5 * time.Second)
}

func compute(msg string)  {
	rand.Seed(time.Now().UnixNano())
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}

}
