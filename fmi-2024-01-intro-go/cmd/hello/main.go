package main

import (
	"fmt"
	"github.com/iproduct/coursego/fmi-2024-01-intro-go/stringutil"
)

func main() {
	s := "Hello Go World - 你好，围棋世界"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
	result := stringutil.Reverse(s)
	fmt.Println(result)
}
