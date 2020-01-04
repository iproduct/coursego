package main

import (
	"strconv"
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Enter a number: ")
	fmt.Scanln(&n)
	str := strconv.FormatInt(int64(n), 10)
	hexadecimal, _ := strconv.ParseInt(str, 16, 64)
	fmt.Printf("%d\n", hexadecimal)
}