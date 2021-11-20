package main

import (
	"fmt"
	"log"
	"os"
)

type Role int

var myname = "Trayan"

func main() {
	fmt.Println(os.Args)
	var pi uint64 = 3
	fmt.Println(pi)
	myNameLocal := "Trayan2"
	fmt.Println(myNameLocal)

	fmt.Println()
	for i := 0x01F500; i < 0x01F700; i++ {
		fmt.Print(string(rune(i)))
	}
	fmt.Println()
	for i := 0; i < 60; i++ {
		fmt.Print(string('🟪'))
	}

	fmt.Println()
	s := "abc🔦🔬🔭"
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		log.Printf("%#U starts at byte position %d\n", r[i], i)
	}

	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U starts at byte position %d\n", s[i], i)
	}

}
