package stringutil

import (
	"fmt"
)

func ExampleReverse() {
	reversed := Reverse("queue")
	fmt.Println(reversed)
	// Output: eueuq
}
