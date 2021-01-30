package main

import "fmt"

type T0 struct {
	x  int
	xp *int
}

func (*T0) M0() int { return 1 }

type T1 struct {
	y int
}

func (T1) M1() int  { return 1 }
func (*T1) M4() int { return 1 }

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2() int { return 1 }

type Q *T2

type R struct {
	*T2
}

var t T2 = T2{5, T1{6}, &T0{x: 4}} // with t.T0 != nil
var xpv = 42
var p *T2 = &T2{5, T1{6}, &T0{x: 4, xp: &xpv}} // with p != nil and (*p).T0 != nil
var q Q = p
var r R = R{p}

func main() {
	fmt.Println(t.z) // t.z
	fmt.Println(t.x) // (*t.T0).x
	fmt.Println(t.y) // t.T1.y

	fmt.Println(p.z) // (*p).z
	fmt.Println(p.y) // (*p).T1.y
	fmt.Println(p.x) // (*(*p).T0).x

	fmt.Println(q.x)            // (*(*q).T0).x        (*q).x is a valid field selector
	fmt.Println(*q.xp)          // (*(*q).T0)).xp is a valid field selector
	fmt.Println(*(*(*q).T0).xp) // (*(*q).T0).xp is a valid field selector

	fmt.Println(p.M0())    // ((*p).T0).M0()      M0 expects *T0 receiver
	fmt.Println(p.M1())    // ((*p).T1).M1()      M1 expects T1 receiver
	fmt.Println(p.M2())    // p.M2()              M2 expects *T2 receiver
	fmt.Println(t.M2())    // (&t).M2()           M2 expects *T2 receiver
	fmt.Println(t.M1())    // (&t).M1()           M1 expects T1 receiver
	fmt.Println(t.M4())    // (&t).M4()           M4 expects *T1 receiver
	fmt.Println(t.M0())    // (&t).M0()           M0 expects *T0 receiver
	fmt.Println((*q).M0()) // (*q).M0()           M0 expects *T0 receiver
	//fmt.Println(q.M0())    // (*q).M0 is valid but not a field selector
	fmt.Println(r.M0()) // (*q).M0 is valid but not a field selector

}
