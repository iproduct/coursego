package main

import (
	"fmt"
	"time"
)

func main() {
	orderReady := make(chan struct{})
	go func() {
		fmt.Println("Chef is preparing the pizza ...")
		time.Sleep(3 * time.Second)
		fmt.Println("Pizza is ready!")
		orderReady <- struct{} {}
	}()
	fmt.Println("Taking orders from clients")
	<- orderReady
	fmt.Println("Pizza is served ...")
	fmt.Println("Simulation complete.")
}
