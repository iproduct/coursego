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
	myIntContainer := &MyContainer{}
	myIntContainer.Put(12.0)
	myIntContainer.Put(70)
	elem, ok := myIntContainer.Get().(int)
	if !ok {
		fmt.Println("Error getting int from myIntContainer")
	}
	fmt.Printf("Got: %d (%T)\n", elem, elem)
}
