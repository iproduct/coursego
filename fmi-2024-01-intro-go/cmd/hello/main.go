package main

import (
	"fmt"
	"github.com/iproduct/coursego/fmi-2024-01-intro-go/stringutil"
)

func main() {
	s := "Hello Go World"
	result := stringutil.Reverse(s)
	fmt.Println(result)
}
