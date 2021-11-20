package main

import "fmt"

func task(i *int) {
	// compute smthng ...
	*i++
	fmt.Println(*i)
}

func main() {

	condition := func(i int) bool {
		return i%5 != 0
	}

	//do {
	//	task();
	//} while (condition);

	// 1)
	i := 0
	for ok := true; ok; ok = condition(i) {
		task(&i)
	}

	// 2)
	fmt.Println()
	i = 0
	for {
		task(&i)
		if !condition(i) {
			break
		}
	}


}
