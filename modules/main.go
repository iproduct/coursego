package main

import (
	"fmt"
	// "github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("Starting the REST Demo")
	a := App{}
	a.Init("root", "root", "go_rest_api")
	a.Run(":8080")
}
