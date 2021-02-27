package main

import (
	"fmt"
	maps "github.com/iproduct/coursego/13-generation-lab/genny/concurrent-map"
	"github.com/iproduct/coursego/13-generation-lab/genny/queue"
)

func main() {
	intQueue := queue.NewIntQueue()
	intQueue.Push(12)
	intQueue.Push(9)
	intQueue.Push(2)
	intQueue.Push(54)

	fmt.Println(intQueue.Pop())
	fmt.Println(intQueue.Pop())
	fmt.Println(intQueue.Pop())
	fmt.Println(intQueue.Pop())

	stringQueue := queue.NewStringQueue()
	stringQueue.Push("abc")
	stringQueue.Push("deef")
	stringQueue.Push("fgh")
	stringQueue.Push("xyz")

	fmt.Println(stringQueue.Pop())
	fmt.Println(stringQueue.Pop())
	fmt.Println(stringQueue.Pop())
	fmt.Println(stringQueue.Pop())

	todos := map[int]string{
		0: "Learn how to use genny",
		1: "Write concurrent map code",
		2: "Generate the typesafe concurrent map implemetations",
	}

	concurrentTodos := maps.ToConMapIntString(todos)

	for k := 0;  k < concurrentTodos.Len(); k++ {
		fmt.Printf("%v --> %v\n", k, concurrentTodos.Get(k))
	}

}
