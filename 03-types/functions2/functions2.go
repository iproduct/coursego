package main

import (
	"fmt"
)

func main() {
	val := "ABCD"
	fmt.Printf("The value before function call is %v\n", val)
	changeValue(val)       // does not change value
	changeValueByRef(&val) //changes value
	fmt.Printf("The value after function call is %v\n", val)
}

func changeValue(num string) {
	num = "NEW_VALUE"
}

func changeValueByRef(num *string) {
	*num = "NEW_VALUE"
}
