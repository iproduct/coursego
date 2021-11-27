package main

import (
	"fmt"
	"reflect"
)

func printSlice(message string, x interface{}) {
	switch y := x.(type) {
	case []int:
		fmt.Printf("%s len=%d cap=%d. %#v\n", message, len(y), cap(y), y)
	case []string:
		fmt.Printf("%s len=%d cap=%d. %#v\n", message, len(y), cap(y), y)
	case []interface{}:
		fmt.Printf("%s len=%d cap=%d. %#v\n", message, len(y), cap(y), y)
	}
}

func printSliceReflection(message string, slice interface{}) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}
	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	fmt.Printf("%s len=%d cap=%d. %#v\n", message, len(ret), cap(ret), ret)
}

//func printStringSlice(message string, x []string) {
//	fmt.Printf("%s len=%d cap=%d. %#v\n", message, len(x), cap(x), x)
//}
//func printIntSlice(message string, x []int) {
//	fmt.Printf("%s len=%d cap=%d. %#v\n", message, len(x), cap(x), x)
//}

func main() {
	a := [...]string{"Hello", "Golang", "!"}
	//var sa1 []interface{}
	sa1 := a[2:len(a)]
	printSlice("sa1 is ", sa1)

	b := make([]int, 5, 10)
	printSlice("b is ", b)

	c := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	printSlice("c is ", c)
	sc1 := c[2:5:7]
	sc2 := sc1[2:3:3]
	printSlice("sc1 is ", sc1)
	printSlice("sc2 is ", sc2)

	sc3 := append(sc1, 11, 12)
	printSlice("c is ", c)
	printSlice("sc3 is ", sc3)

	var sc4 []interface{} // zero value == nil
	for _, v := range c {
		sc4 = append(sc4, v)
	}
	printSlice("sc4", sc4)

	sc5 := make([]int, 5)
	copy(sc5, c)
	printSlice("sc5", sc5)
}
