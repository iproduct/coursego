// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mypic // import "github.com/iproduct/coursego/03-types-lab/mypic"

import (
	"image"
	"image/png"
	"os"
)

// Encode shows the image provided by the factory function f and writes it to provided PNG file
func Encode(f func(int, int) [][]uint8, file *os.File) {
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
	EncodeImagePNG(m, file)
}

// EncodeImagePNG encodes the image and writes it to PNG file
func EncodeImagePNG(m image.Image, file *os.File) error {
	// var buf bytes.Buffer
	// err := png.Encode(&buf, m)
	// if err != nil {
	// 	panic(err)
	// }
	// enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	// fmt.Fprintln(file, "IMAGE:" + enc)
	return png.Encode(file, m)
}
