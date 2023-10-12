package main

import "fmt"

func main() {
	s := "abc日本語"
	for index, rune := range s {
		fmt.Printf("%#U starting at %d\n", rune, index)
	}
	fmt.Println()
	for index := 0; index < len(s); index++ {
		fmt.Printf("%#U starting at %d\n", s[index], index)
	}
	fmt.Println()
	runes := []rune(s)
	for index := 0; index < len(runes); index++ {
		fmt.Printf("%#U starting at %d\n", runes[index], index)
	}
	runes[0] = 'z'
	fmt.Println(s)
	fmt.Println(string(runes))

}
