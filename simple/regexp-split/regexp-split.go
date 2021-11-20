package main

import (
	"fmt"
	"regexp"
)

func main()  {
	a := regexp.MustCompile(`n`)              // a single `a`
	fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
	fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
	fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
	fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

	zp := regexp.MustCompile(` *, *`)             // spaces and one comma
	fmt.Printf("%q\n", zp.Split("a,b ,  c ", -1)) // ["a" "b" "c "]
}
