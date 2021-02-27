package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func SortPeople(people []Person) {
	sort.Sort(ByAge(people))
}
func main() {
	people := []Person{
		{"Bob", 75},
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
	}
	SortPeople(people)
	fmt.Println("By age:", people)
}
