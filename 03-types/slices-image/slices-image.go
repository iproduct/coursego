package main

import (
	"fmt"
	"github.com/iproduct/coursego/03-types-lab/mypic"
	"log"
	"os"
	"path"
)

const baseDir = "d:/CourseGO/git/coursegopro/03-types/slices-image"

// Pic returns a grayscale pic of size dy * dx
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
		for j := range result[i] {
			result[i][j] = uint8(i ^ j)
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
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mypic.Encode(Pic, file)
}
