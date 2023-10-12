package main

import "fmt"

func zeroval(ival int)  { ival = 0 }
func zeroptr(iptr *int) { *iptr = 0 }

func main() {
	//n := 42
	//var p = &n
	//*p++
	//fmt.Printf("%p -> %d , n = %d\n", p, *p, n)

	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
