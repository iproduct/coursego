package main

import "fmt"

func main() {
	var s uint = 33
	var i = 1<<s                  // 1 has type int
	fmt.Printf("%#v\n", i)
	var j int32 = 1<<s            // 1 has type int32; j == 0
	fmt.Println(j)
	var k = uint64(1<<s)          // 1 has type uint64; k == 1<<33
	fmt.Println(k)
	var m int = 1.0<<s            // 1.0 has type int; m == 0 if ints are 32bits in size
	fmt.Println(m)
	var n = (int32(1.0)<<s) == j           // 1.0 has type int32; n == true
	fmt.Println(n)
	var o = 1<<s == 2<<s          // 1 and 2 have type int; o == true if ints are 32bits in size, false otherwise
	fmt.Println(o)
	//var p = 1<<s == 1<<64         // illegal if ints are 32bits in size: 1 has type int, but 1<<33 overflows int
	//fmt.Println(p)
	//var u = 1.0<<s                // illegal: 1.0 has type float64, cannot shift
	//fmt.Println(u)
	//var u1 = 1.0<<s != 0          // illegal: 1.0 has type float64, cannot shift
	//fmt.Println()
	//var u2 = 1<<s != 1.0          // illegal: 1 has type float64, cannot shift
	//fmt.Println(u2)
	//var v float32 = 1<<s          // illegal: 1 has type float32, cannot shift
	//fmt.Println()
	var w int64 = 1.0<<33         // 1.0<<33 is a constant shift expression
	fmt.Println(w)
	a := []int{12, 15, 17}
	var x = a[1.0<<1]             // 1.0 has type int; x == a[2] if ints are 32bits in size
	fmt.Println(x)
	var b = make([]byte, 1.0<<3)  // 1.0 has type int; len(a) == 0 if ints are 32bits in size
	fmt.Printf("%#v -> len: %d, cap: %d", b, len(b), cap(b))
}
