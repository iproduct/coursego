package main

import "fmt"

func main() {
	const a = 2 + 3.0        // a == 5.0   (untyped floating-point constant)
	const b = 15 / 4         // b == 3     (untyped integer constant)
	const c = 15 / 4.0       // c == 3.75  (untyped floating-point constant)
	const Θ float64 = 3 / 2  // Θ == 1.0   (type float64, 3/2 is integer division)
	const Π float64 = 3 / 2. // Π == 1.5   (type float64, 3/2. is float division)
	const d = 1 << 3.0       // d == 8     (untyped integer constant)
	fmt.Printf("%v : %[1]T\n", d)
	const e = 1.0 << 3 // e == 8     (untyped integer constant)
	//const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
	//const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
	const h = "foo" > "bar" // h == true  (untyped boolean constant)
	const j = true          // j == true  (untyped boolean constant)
	const k = 'w' + 1       // k == 'x'   (untyped rune constant)
	fmt.Printf("%c : %[1]U: %[1]T\n", k)
	const l = "hi"         // l == "hi"  (untyped string constant)
	const m = string(k)    // m == "x"   (type string)
	const Σ = 1 - 0.707i   //            (untyped complex constant)
	const Δ = Σ + 2.0e-4   //            (untyped complex constant)
	const Φ = iota*1i - 1  //            (untyped complex constant)
	const Φ2 = iota*1i - 1 //            (untyped complex constant)
	fmt.Printf("%f : %[1]T\n", Φ)
	fmt.Printf("%f : %[1]T\n", Φ2)

	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 // untyped floating-point constant
	const (
		size int64 = 1024
		eof        = -1 // untyped integer constant
	)
	const x, y, z = 3, 4, "foo" // a = 3, b = 4, c = "foo", untyped integer and string constants
	const u, v float32 = 0, 3   // u = 0.0, v = 3.0

	const (
		c1 = imag(2i)                   // imag(2i) = 2.0 is a constant
		c2 = len([10]float64{2})        // [10]float64{2} contains no function calls
		c3 = len([10]float64{c1})       // [10]float64{c1} contains no function calls
		c4 = len([10]float64{imag(2i)}) // imag(2i) is a constant and no function call is issued
		//c5 = len([10]float64{imag(z)})   // invalid: imag(z) is a (non-constant) function call
	)
	comp := imag(3 + 2i)
	c5 := [10]float64{comp}
	fmt.Printf("%#v\n", c5)

	const Huge = 1 << 100        // Huge == 1267650600228229401496703205376  (untyped integer constant)
	const Four int8 = Huge >> 98 // Four == 4                                (type int8)
	fmt.Printf("%#v\n", Four)
	const (
	//c21 = uint(-1)     // -1 cannot be represented as a uint
	//c22 = int(3.14)    // 3.14 cannot be represented as an int
	//c23 = int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
	//c24= Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
	//c25 = Four * 100   // product 400 cannot be represented as an int8 (type of Four)
	)
	//fmt.Printf("%#v\n", c22)
}
