package main

import (
	"fmt"
	"github.com/iproduct/coursego/03-types-lab/mypic"
	"log"
	"os"
	"path"
)

const baseDir = "d:/CourseGO/workspace/src/github.com/iproduct/coursego/03-types-lab/slices-image"

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

	imageFile := path.Join(baseDir, "image.png")
	fmt.Println(imageFile)
	file, err := os.Create(imageFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	mypic.Encode(Pic, file)
}
