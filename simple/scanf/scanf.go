package main

import "fmt"

func main () {
	count:= 0
	for {
		var input int
		fmt.Scanf ("% v \n",&input)
		count ++
		fmt.Printf ("input:% v \n", input)
		fmt.Printf ("count:% v \n", count)
		if count == 10 {
			break
		}
	}
}