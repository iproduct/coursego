package main

import (
	"fmt"
	strutil "github.com/iproduct/coursego/fmi-2025-01-intro/stringutil"
)

func main() {
	s := "Hello Go World - 你好，围棋世界!"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%c", r)
	}
	fmt.Println()
	result := strutil.Reverse(s)
	fmt.Println(result)
}
