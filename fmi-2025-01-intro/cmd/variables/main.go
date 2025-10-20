package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var global int64

var (
	user   = os.Getenv("USER")
	home   = os.Getenv("HOME")
	gopath = os.Getenv("GOPATH")
)

func countLinesFile(f *os.File) (int, error) {
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		//log.Println(scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return count, nil
}

func CountLines(filename string) int {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	count, err := countLinesFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func main() {
	global = 5
	fmt.Printf("%v, %v, %v\n", user, home, gopath)
	fmt.Printf("%[1]v, %[1]T\n", global)
	cp1, cp2 := 12+3i, 7+5i
	cp3 := cp1 + cp2
	fmt.Println(cp3)
	pcp3 := &cp3
	fmt.Printf("%[1]p, %[1]T\n", pcp3)
	pint := new(int)
	fmt.Printf("%[1]v, %[1]T\n", *pint)
	channel := make(chan interface{})
	go func(size int64) {
		defer close(channel)
		var i int64
		for i = 0; i < size; i++ {
			channel <- "message_" + strconv.Itoa(int(i))
			time.Sleep(time.Second)
		}
	}(global)

	//for ok := true; ok; {
	//	var result interface{}
	//	result, ok = <-chan2
	//	if ok {
	//		fmt.Println(result)
	//	} else {
	//		fmt.Println("channel closed")
	//	}
	//}

	chan2 := (<-chan interface{})(channel)
	for result := range chan2 {
		fmt.Println(result.(string))
	}

	// Count lines in own sorce file
	fmt.Printf("Number of  lines: %v\n", CountLines("cmd/variables/main.go"))
}
