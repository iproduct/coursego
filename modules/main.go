package main

import (
	"fmt"
	"github.com/iproduct/coursego/modules/rest"
	// "github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("Starting the REST Demo")
	a := rest.App{}
	a.Init("root", "root", "go_rest_api")
	a.Run(":8080")
}
