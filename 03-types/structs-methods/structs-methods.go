package main

import "fmt"

type T0 struct {
	x int
}

func (*T0) M0() int { return 1 }

type T1 struct {
	y int
}

func (T1) M1() int { return 1 }

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2() int { return 1 }

type Q *T2

var t T2 = T2{5, T1{6}, &T0{4}}   // with t.T0 != nil
var p *T2 = &T2{5, T1{6}, &T0{4}} // with p != nil and (*p).T0 != nil
var q Q = p

func main() {
	fmt.Println(t.z) // t.z
	fmt.Println(t.x) // (*t.T0).x
	fmt.Println(t.y) // t.T1.y

	fmt.Println(p.z) // (*p).z
	fmt.Println(p.y) // (*p).T1.y
	fmt.Println(p.x) // (*(*p).T0).x

	fmt.Println(q.x) // (*(*q).T0).x        (*q).x is a valid field selector

	fmt.Println(p.M0())    // ((*p).T0).M0()      M0 expects *T0 receiver
	fmt.Println(p.M1())    // ((*p).T1).M1()      M1 expects T1 receiver
	fmt.Println(p.M2())    // p.M2()              M2 expects *T2 receiver
	fmt.Println(t.M2())    // (&t).M2()           M2 expects *T2 receiver, see section on Calls
	fmt.Println(t.M1())    // (&t).M1()           M2 expects *T2 receiver, see section on Calls
	fmt.Println(t.M0())    // (&t).M0()           M2 expects *T2 receiver, see section on Calls
	fmt.Println((*q).M0()) // (&t).M2()           M2 expects *T2 receiver, see section on Calls
}
