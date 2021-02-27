package main

import "fmt"

//go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	//VitaminC
	Paracetamol
	Acetaminophen = Paracetamol
)

func main() {
	fmt.Println(Acetaminophen)
	//fmt.Println(VitaminC)
}
