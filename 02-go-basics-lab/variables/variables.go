package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
)

var myglobal float32 = 3.14

func countLines(f *os.File) int {
	input := bufio.NewScanner(f)
	var count int
	for input.Scan() {
		count++
	}
	return count
}

func Sample(name string) (int, error) {
	fp, err := os.Open(name)
	//defer func() {
	//	err = fp.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	if err != nil {
		return 0, err
	}
	defer fp.Close()
	return countLines(fp), nil
}

var (
	gl1 = 42
	gl2 = 108
)

func main() {
	var local1 int16 = 35
	local2 := 12000
	fmt.Printf("myvalue = %#V\n", myglobal)
	fmt.Printf("gl1 = %#V\n", gl1)
	fmt.Printf("gl2 = %#V\n", gl2)
	fmt.Printf("local1 = %#V\n", local1)
	fmt.Printf("local2 = %#V\n", local2)
	numlines, err := Sample(`D:\CourseGO\git\coursego\02-go-basics-lab\variables\variables.go`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number lines in program: %d\n", numlines)
}