package main

import (
	"fmt"
	"reflect"
)

type Stack struct {
	t          reflect.Type
	collection reflect.Value
}

func New(tp reflect.Type) *Stack {
	return &Stack{
		t:          tp,
		collection: reflect.MakeSlice(reflect.SliceOf(tp), 0, 50),
	}
}

func (s *Stack) Push(elem interface{}) {
	if reflect.ValueOf(elem).Type() != s.t {
		panic(fmt.Sprintf("Error putting %T into stack of type %s", elem, s.t))
	}
	s.collection = reflect.Append(s.collection, reflect.ValueOf(elem))
}

func (s *Stack) Pop() interface{} {
	v := s.collection.Index(s.collection.Len()-1)
	s.collection = s.collection.Slice(0, s.collection.Len()-1)
	return v.Interface()
}

func main() {
	v := "12.5"
	m := New(reflect.TypeOf(v))
	m.Push(v)
	m.Push("82.13")

	for i := 0; i < 2; i++ {
		elem:= m.Pop()
		fmt.Printf("Result: %[1]T, %[1]v \n", elem)
	}
}
