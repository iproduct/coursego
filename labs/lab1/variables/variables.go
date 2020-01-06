package main

import (
	"flag"
	"log"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"os"
)

var global int = 50

var (
    gopath = os.Getenv("GOPATH")
)

func init() {
	global = 12
   
    // gopath may be overridden by --gopath flag on command line.
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")

	if gopath == "" {
		gopath = "c:/coursego/workspace"
		log.Printf("GOPATH not set - using default: %s", gopath)
   }
}
func sample(name string ) error {
	var i, j int = 5, 9
	i, j = j, i // swap values of i and j
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	// ...use f...
	f.Close()
	return nil
}

func main() {
	fmt.Printf("GOPATH=%v\n", gopath)
	// global = 42
	// var local int = 5
	local, n := 5, 12
	fmt.Printf("global = %d\nlocal = %d\n", global, local+n)
	var cp1 complex128 = 12 + 3i
	cp2 := 5 + 7i
	fmt.Printf("%#v\n", cp1+cp2)
	str := "Hello"
	fmt.Printf("%-20.20s -> %5d\n", str, n)

	chA := make(chan interface{})
	chB := make(chan interface{})
	go func() {
		defer close(chA)
		var i int64
		for i = 0; i < 10; i++ {
			chA <- "A_message_" + strconv.FormatInt(i, 10)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	go func() {
		defer close(chB)
		var i int64
		for i = 0; i < 10; i++ {
			chB <- "B_message_" + strconv.FormatInt(i, 10)
			time.Sleep(time.Second)
		}
	}()

	toComplete := 2
	for {

		select {
		case result, ok := <-chA:
			if ok {
				fmt.Printf("Channel A: %v\n", result)
			} else {
				chA = nil
				toComplete--
			}

		case result, ok := <-chB:
			if ok {
				fmt.Printf("Channel B: %v\n", result)
			} else {
				chB = nil
				toComplete--
			}
		}

		if toComplete == 0 {
			fmt.Printf("END\n")
			break
		}
	}
}
