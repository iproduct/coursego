package main

import (
	"fmt"
	"strings"
)


func main() {
	wordsChannel := LengthGreaterThan4(Capitalize(ProduceEvents()))
	for event := range wordsChannel {
		fmt.Println(event) // => HELLO, REACTIVE, WORLD
	}
}


func ProduceEvents() <-chan string {
	words := []string {"Hello", "Reactive", "World", "from", "RxGo"}
	out := make(chan string)
	go func() {
		for _, word := range words {
			out <- word
			//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(out)
	}()
	return out
}

func LengthGreaterThan4(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for item := range in {
			if len(item) > 4 {
				out <- item
			}
		}
		close(out)
	}()
	return out
}

func Capitalize(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for item := range in {
			out <- strings.ToUpper(item)
		}
		close(out)
	}()
	return out
}
