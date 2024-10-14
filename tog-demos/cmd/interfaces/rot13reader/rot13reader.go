package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(buf []byte) (int, error) {
	n, err := r13.r.Read(buf)
	if err == nil {
		for i := 0; i < n; i++ {
			//fmt.Println(i, buf[i])
			if buf[i] >= 'A' && buf[i] <= 'Z' {
				buf[i] = (buf[i]-'A'+13)%26 + 'A'
			} else if buf[i] >= 'a' && buf[i] <= 'z' {
				buf[i] = (buf[i]-'a'+13)%26 + 'a'
			}
		}
	} else {
		fmt.Println(err)
		return n, err
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
