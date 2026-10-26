package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{Lat: 40.68433, Long: -74.39967},
	"Google":    Vertex{37.42202, -122.08408},
}

var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

var s1 = [...]Vertex{
	{40.68433, -74.39967},
	{37.42202, -122.08408},
}
var s2 = [...]Vertex{
	{40.68433, -74.39967},
	{37.42202, -122.08408},
}

func main() {
	fmt.Println(m2)
	fmt.Printf("%T, equal-> %t\n", m2, m["Bell Labs"] == m2["Bell Labs"])
	fmt.Printf("%T, equal-> %t\n", s1, s1 == s2)
}
