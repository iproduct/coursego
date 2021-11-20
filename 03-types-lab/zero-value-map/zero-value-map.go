package main

import "fmt"

type Node struct {
	Next  *Node
	Value interface{}
}

func VisitNodes(head *Node) {
	visited := make(map[*Node]struct{})
	for n := head; n != nil; n = n.Next {
		if _, ok := visited[n]; ok {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = struct {} {}
		fmt.Println(n.Value)
	}
}

func main() {
	node := &Node{nil, "A"}
	end := node
	//var node *Node = nil
	for i := range make([]struct{}, 5){
		node = &Node{node, string('Z' - i)}
	}
	end.Next = node
	VisitNodes(node)

}
