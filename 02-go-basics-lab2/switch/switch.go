package main

import (
	"fmt"
	"time"
)

func main() {
	pred := func(hour int) bool {
		afterHours := []int{12, 13, 14, 15, 16, 17}
		for _, v := range afterHours {
			if v == hour {
				return true
			}
		}
		return false
	}
	switch h := time.Now().Hour(); {
	case pred(h):
		fmt.Println("Good afternoon.")
	case h > 17:
		fmt.Println("Good evening.")
	default:
		fmt.Println("Good morning!")
	}

}
