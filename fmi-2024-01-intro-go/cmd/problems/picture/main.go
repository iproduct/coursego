package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func Pic(dx, dy int) [][]uint8 {
	return make([][]uint8, dy) //TODO draw the picture
}

func main() {
	f, err := os.OpenFile("image.png", os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	SaveSliceToFile(Pic, f)
}

// SaveImageToFile saves the slice m
// to given file in PNG format.
func SaveSliceToFile(f func(dx, dy int) [][]uint8, file *os.File) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(file, m)
	if err != nil {
		panic(err)
	}
}
