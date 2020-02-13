package main

import "fmt"

type MyContainer []interface{}

func (c *MyContainer) Put(elem interface{}) {
	*c = append(*c, elem)
}

func (c *MyContainer) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}

func main() {
	m := &MyContainer{}
	m.Put(12)
	m.Put(42)

	for i := 0; i < 2; i++ {
		elem, ok := m.Get().(int)
		if !ok {
			fmt.Println("Error getting element - not int.")
		}
		fmt.Println(elem)
	}
}
