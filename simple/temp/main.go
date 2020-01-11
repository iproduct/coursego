package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"time"
	"unsafe"
)

// Sequence represents a sequence of integers
type Sequence []int

// Copy method copies the current to a new Sequence
func (s Sequence) Copy() Sequence {
	result := make([]int, len(s))
	copy(result, s)
	return result
}

// String method sorts elements and returns them as string
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

// Get returns the i-th element of the Sequence
func (s Sequence) Get(i int) int {
	return s[i]
}

// Stringer interface declares single method String() string
type Stringer interface {
	String() string
}

func myString(value interface{}) (string, error) {
	switch str := value.(type) {
	case string:
		return str, nil
	case Stringer:
		return str.String(), nil
	}

	if str, ok := value.(string); ok {
		return str, nil
	} else if str, ok := value.(Stringer); ok {
		return str.String(), nil
	} else {
		return "", errors.New("Type not recongized")
	}
}

type mystringer struct {
	value string
}

func (ms mystringer) String() string {
	return ms.value
}

func loops() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func loopForMissing() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func loopsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func loopsDoWhile() {
	sum := 1
	for {
		sum += sum
		if sum > 1000 {
			break
		}
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func switch1() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fallthrough
	case today + 3:
		fmt.Println("Have to wait some time")
	default:
		fmt.Println("Too far away.")
	}
}

func switch2() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

//func swap(x, y string) (string, string) {
//	return y, x
//}
//
//func main() {
//	a, b := swap("hello", "world")
//	fmt.Println(a, b)
//}
//
//func split(sum int) (x, y int) {
//	x = sum * 4 / 9
//	y = sum - x
//	return
//}
//
//func main() {
//	fmt.Println(split(17))
//}

func strToBytes(s string) []byte {
	header := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytesHeader := &reflect.SliceHeader{
		Data: header.Data,
		Len: header.Len,
		Cap: header.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bytesHeader))
}

func bytesToStr(b []byte) string  {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	stringHeader := &reflect.StringHeader{
		Data: header.Data,
		Len: header.Len,
	}
	return *(*string)(unsafe.Pointer(stringHeader))
}

func swapVal(x, y string) (string, string) {
	return y, x
}
func swapRef(x, y *string) {
	*x, *y = *y, *x
}

func main() {
	a, b := swapVal("hello", "world")
	fmt.Println(a, b)
	swapRef(&a, &b)
	fmt.Println(a, b)
}

//func main() {
//	fmt.Println(myString(mystringer{"abcd"}))
//	fmt.Println(myString("xyz"))
//	s := Sequence{1, 5, 7, 2, 9, 3}
//	fmt.Println(myString(s))
//	fmt.Println(s.Get(1))
//	loops()
//	switch1()
//
//	a := "hello"
//	header := (*reflect.StringHeader)(unsafe.Pointer(&a))
//	header.Len = 100
//	// cast the header back to 'string' and print it
//	fmt.Println(*(*string)(unsafe.Pointer(header)))
//	//v := reflect.ValueOf(a)
//	fmt.Printf("%#v\n", header)
//	//s := reflect.ValueOf(&t).Elem()
//	//typeOfA := v.Type()
//	//for i := 0; i < v.String(); i++ {
//	//	f := v.Field(i)
//	//	fmt.Printf("%d: %s %s = %v\n", i,
//	//		typeOfA.Field(i).Name, f.Type(), f.Interface())
//	//}
//	a = a + " there"
//	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&a)))
//	a = "world123"
//	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&a)))
//
//	//b := strToBytes("hello")
//	//b[0] = 100
//	//fmt.Print(b)
//
//	b := []byte{104, 101, 108, 108, 111}
//	str := bytesToStr(b)
//	fmt.Println(str) // "hello"
//	b[0] = 100
//	fmt.Println(str) // "dello"
//}
