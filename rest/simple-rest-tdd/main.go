// main.go

package main

func main() {
	a := App{}
	a.Initialize("root", "root", "go_rest_api")

	a.Run(":8080")
}
