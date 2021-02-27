package main

import (
	"fmt"
	"github.com/iproduct/coursego/generationlab/genny/maps"
	"github.com/iproduct/coursego/generationlab/genny/queue"
)

func main(){
	q := queue.NewStringQueue()
	q.Push("abc")
	q.Push("def")
	q.Push("xyz")
	q.Push("1234")

	for i:= 0; i< 4 ; i++ {
		fmt.Println(q.Pop())
	}

	iq := queue.NewIntQueue()
	iq.Push(1)
	iq.Push(2)
	iq.Push(3)
	iq.Push(42)

	for i:= 0; i< 4 ; i++ {
		fmt.Println(iq.Pop())
	}

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