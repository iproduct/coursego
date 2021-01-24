package main

import "fmt"

type Node struct {
	Next  *Node
	Value interface{}
}

func VisitNodes(head *Node) {
	visited := make(map[*Node]bool)
	for n := head; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}
}

func main() {

}
