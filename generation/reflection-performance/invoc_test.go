package main

import (
	"reflect"
	"testing"
)

type myint int64

type Inccer interface {
	inc()
}

func (i *myint) inc() {
	*i = *i + 1
}

func BenchmarkReflectMethodCall(b *testing.B) {
	i := new(myint)
	incnReflectCall(i.inc, b.N)
}

func BenchmarkReflectOnceMethodCall(b *testing.B) {
	i := new(myint)
	incnReflectOnceCall(i.inc, b.N)
}

func BenchmarkStructMethodCall(b *testing.B) {
	i := new(myint)
	incnIntmethod(i, b.N)
}

func BenchmarkInterfaceMethodCall(b *testing.B) {
	i := new(myint)
	incnInterface(i, b.N)
}

func BenchmarkTypeSwitchMethodCall(b *testing.B) {
	i := new(myint)
	incnSwitch(i, b.N)
}

func BenchmarkTypeAssertionMethodCall(b *testing.B) {
	i := new(myint)
	incnAssertion(i, b.N)
}

func incnReflectCall(v interface{}, n int) {
	for k := 0; k < n; k++ {
		reflect.ValueOf(v).Call(nil)
	}
}

func incnReflectOnceCall(v interface{}, n int) {
	fn := reflect.ValueOf(v)
	for k := 0; k < n; k++ {
		fn.Call(nil)
	}
}

func incnIntmethod(i *myint, n int) {
	for k := 0; k < n; k++ {
		i.inc()
	}
}

func incnInterface(any Inccer, n int) {
	for k := 0; k < n; k++ {
		any.inc()
	}
}

func incnSwitch(any Inccer, n int) {
	for k := 0; k < n; k++ {
		switch v := any.(type) {
		case *myint:
			v.inc()
		}
	}
}

func incnAssertion(any Inccer, n int) {
	for k := 0; k < n; k++ {
		if newint, ok := any.(*myint); ok {
			newint.inc()
		}
	}
}
