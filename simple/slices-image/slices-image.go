package main

import (
	"fmt"
	"github.com/iproduct/coursego/simple/mypic"
	"log"
	"os"
	"path/filepath"
)

const baseDir = "d:/CourseGO/workspace/src/github.com/iproduct/coursego/simple/slices-image"

// Pic returns a grayscale pic of size dy * dx
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		result[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			result[y][x] = uint8(x * y)
		}
	}
	return result
}

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
	// programPath := "os.Args[0]"
	// fmt.Println(programPath)
	// dir := path.Dir(programPath)
	// fmt.Println(dir)

	baseDir := "./images" // Relative path works fine!
	filename := "image.png"

	// 1. Get the absolute path of your intended base directory
	absBase, err := filepath.Abs(baseDir)
	if err != nil {
		panic(err)
	}
	// 2. Join the path
	imageFile := filepath.Join(absBase, filename)

	fmt.Println(imageFile)
	file, err := os.Create(imageFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	mypic.Encode(Pic, file)
}
