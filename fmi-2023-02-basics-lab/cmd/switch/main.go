package main

import (
	"fmt"
	"time"
)

func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

func main() {
	// 1 with default
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Today is week day")
	}

	//2 no condition
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// 3 case list
	fmt.Printf("\\t is white space: %v\n", WhiteSpace('\t'))

	// 4 fallthrough
	switch 1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}

	// 5 exit with break
Loop:
	for _, ch := range "a b \nc" {
		switch ch {
		case ' ':
			break
		case '\n':
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}
}
