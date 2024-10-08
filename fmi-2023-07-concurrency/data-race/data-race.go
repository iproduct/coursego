package main

import (
	"fmt"
	"time"
)

type RPC struct {
	result int
	done   chan struct{}
}

func (rpc *RPC) compute() {
	time.Sleep(time.Second) // strenuous computation intensifies
	rpc.result = 42
	//rpc.done <- struct{}{}
	close(rpc.done)
}

func (RPC) version() int {
	return 1 // never going to need to change this
}

func main() {
	rpc := &RPC{done: make(chan struct{})}

	go rpc.compute()         // kick off computation in the background
	version := rpc.version() // grab some other information while we're waiting
	val, ok := <-rpc.done    // wait for computation to finish
	fmt.Printf("%#v, %T, ok = %#v\n", val, val, ok)
	result := rpc.result

	fmt.Printf("RPC computation complete, result: %d, version: %d\n", result, version)
}
