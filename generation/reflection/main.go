package main

import (
	"fmt"
	"reflect"
)

type MyStack struct {
	t reflect.Type
	value reflect.Value
}

func New(tp reflect.Type) *MyStack {
	return &MyStack{
		t: tp,
		value: reflect.MakeSlice(reflect.SliceOf(tp), 0, 100),
	}
}

func (m *MyStack) Push(v interface{}) {
	if reflect.ValueOf(v).Type() != m.value.Type().Elem() {
		panic(fmt.Sprintf("Error putting %T into a container of type %s", v, m.value.Type().Elem()))
	}

	m.value = reflect.Append(m.value, reflect.ValueOf(v))
}

func (m *MyStack) Pop() interface{} {
	v := m.value.Index(0)
	m.value = m.value.Slice(1, m.value.Len())
	return v.Interface()
}

func main() {
	val := 2.88
	fmt.Println(reflect.TypeOf(val))
	stack := New(reflect.TypeOf(val))
	stack.Push(val)
	fmt.Println(stack.value.Index(0))
	result := stack.Pop()
	fmt.Printf("Result: %[1]f (%[1]T)\n", result)
}
