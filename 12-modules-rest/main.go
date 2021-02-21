package main

import (
	"fmt"
	"github.com/iproduct/coursego/moduleslab/rest"
)

func main() {
	fmt.Println("Staring REST User Service ...")
	a := rest.App{}
	a.Init("root", "root", "go_rest_api")
	a.Run(":8080")
}

