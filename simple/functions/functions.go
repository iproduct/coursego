package main

import (
	"fmt"
)

func printf(format string, args ...interface{}) int {
	fmt.Println(args...)
	return len(args)
}

func main() {
	argsLen := printf("%v, %v\n", "abcd", 15, []int{1, 2, 3, 4})
	printf("Number args: %d\n", argsLen)

	// closure demo
	//count := 0
	//inc := func() int {
	//	count++
	//	return count
	//}
	//
	//incBy := func(n int) int {
	//	count += n
	//	return count
	//}
	//
	//printf("%d\n", inc())
	//printf("%d\n", incBy(10))

}
