package main

import (
	"fmt"
	"github.com/iproduct/coursegopro/10-modules-rest-jwtauth/rest"
)

func main() {
	fmt.Println("Staring REST User Service ...")
	a := rest.App{}
	a.Init("root", "root", "go_rest_api")
	a.Run(":8080")
}

