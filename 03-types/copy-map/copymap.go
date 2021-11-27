package main

import "fmt"

func main() {

	// Creating and initializing a map
	// Using shorthand declaration and
	// using map literals
	originalMap := make(map[string]int)
	originalMap["one"] = 1
	originalMap["two"] = 2
	originalMap["three"] = 3
	originalMap["four"] = 4
	originalMap["five"] = 5
	originalMap["six"] = 6
	originalMap["seven"] = 7
	originalMap["eight"] = 8
	originalMap["nine"] = 9

	// Creating empty map
	CopiedMap := make(map[string]int)

	/* Copy Content from Map1 to Map2*/
	for index, element := range originalMap {
		CopiedMap[index] = element
	}

	for index, element := range CopiedMap {
		fmt.Println(index, "=>", element)
	}
}
