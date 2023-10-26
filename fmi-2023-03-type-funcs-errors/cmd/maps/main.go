package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex, 10)
	fmt.Printf("%+v, len=%d\n", m, len(m))
	m["Bell Labs"] = Vertex{40.68433, 74.39967}
	fmt.Printf("%+v, len=%d\n", m, len(m))
	fmt.Printf("%+v\n", m["Bell"])
}
