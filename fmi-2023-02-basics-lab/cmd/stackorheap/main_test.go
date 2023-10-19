package main

import (
	"testing"
)

var num = 1000

func BenchmarkStackOrHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f()
		//fmt.Printf("%d", *global)
		g()
	}
}
