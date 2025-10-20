package main

import (
	"fmt"
	"github.com/iproduct/coursego/fmi-2025-01-intro/sequence"
)

func main() {
	seq := sequence.Sequence([]int{11, 22, 13, 54, 35, 16, 77, 8})
	fmt.Println(seq)
}
