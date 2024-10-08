package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Image struct {
	width  int
	height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{100, 100}
	f, err := os.OpenFile("image.png", os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	SaveImageToFile(m, f)
}

// SaveImageToFile saves the image m
// to given file in PNG format.
func SaveImageToFile(m image.Image, file *os.File) {
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(file, m)
	if err != nil {
		panic(err)
	}
}
