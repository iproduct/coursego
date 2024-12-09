package main

import "fmt"

func main() {
	var c chan string
	fmt.Println(<-c) // Deadlock
}

func main2() {
	select {}
	fmt.Println("Demo finished")
}
