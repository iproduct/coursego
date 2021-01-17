package main

import (
	"fmt"
)

func main() {
	//var str string
	n, sum := 1, 0
	for ok := true; ok; ok = n > 0 {
		fmt.Printf("\nEnter a number [Enter or 0 for end]: ")
		ch, err := fmt.Scanln(&n)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
		if ch == 0 {
			break
		}
		//n, err := strconv.Atoi(str)
		//if err != nil {
		//	//log.Error(err)
		//	fmt.Printf("Error reading a number: %s\n", err.Error())
		//
		//}
		for i := 1; i <= n; i++ {
			fmt.Printf("%d ", i)
		}
		sum += n
	}
	fmt.Printf("\nSUM = %d\n\n", sum)

	// using break
	sum = 0
	for {
		fmt.Printf("Enter a number [Enter for end]: ")
		fmt.Scanf("%d\n", &n)
		fmt.Printf("n = %d\n", n)
		if n <= 0 {
			break
		}
		sum += n
	}
	fmt.Printf("SUM = %d\n\n", sum)

}
