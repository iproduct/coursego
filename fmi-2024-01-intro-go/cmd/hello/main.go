package main

import (
	"fmi-2024-01-intro-go/stringutil"
	"fmt"
)

func main() {
	s := "Hello Go World - 你好，围棋世界"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n ", s[i])
	}

	for _, r := range s {
		fmt.Printf("%c\n ", r)
	}
	result := stringutil.Reverse(s)
	fmt.Println(result)
}
