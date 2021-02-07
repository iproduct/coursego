package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ByteSlice []byte

func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	return len(data), nil
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// A simple File interface.
//type File interface {
//	Read([]byte) (int, error)
//	Write([]byte) (int, error)
//	Close() error
//}

// Implementation of File interface
type MyFile struct { /*...*/
}

func (p MyFile) Read(b []byte) (n int, err error)  { return /*...*/ }
func (p MyFile) Write(b []byte) (n int, err error) { return /*...*/ }
func (p MyFile) Close() error                      { return nil }

// Another implementation of File interface
type OtherFile struct { /*...*/
}

func (o *OtherFile) Read(b []byte) (n int, err error)  { return /*...*/ }
func (o *OtherFile) Write(b []byte) (n int, err error) { return /*...*/ }
func (o *OtherFile) Close() error                      { return nil }

//package io
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}
type File interface {
	Reader
	Writer
	Closer
}

//type ReadWriter interface {
//	Read(b Buffer) bool
//	Write(b Buffer) bool
//}
//type Locker interface {
//	Lock()
//	Unlock()
//}
//type File interface {
//	ReadWriter  // same as adding the methods of ReadWriter
//	Locker      // same as adding the methods of Locker
//	Close()
//}
//type LockedFile interface {
//	Locker
//	File        // illegal: Lock, Unlock not unique
//	Lock()      // illegal: Lock not unique
//}

// illegal: Bad cannot embed itself
//type Bad interface {
//	Bad
//}
//
//// illegal: Bad1 cannot embed itself using Bad2
//type Bad1 interface {
//	Bad2
//}
//type Bad2 interface {
//	Bad1
//}

// Interfaces embedding
type Bouncer interface {
	Bounce()
}
type Football struct {
	Bouncer
}

type Ball struct {
	Radius   int
	Material string
}

func (b Ball) Bounce() {
	fmt.Printf("Ball bouncing ... %v\n", b)
}

// Declaring interface implementation
var _ io.Writer = (*bytes.Buffer)(nil)

func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "Interfaces in Go are %s\n", "useful")
	fmt.Printf("%s\n", b)
	var c ByteCounter
	fmt.Fprintf(&c, "Interfaces in Go are %s\n", "useful")
	fmt.Printf("%v\n", c)

	var f File
	f = MyFile{} //or &MyFile{}
	fmt.Printf("%T\n", f.Close())
	f = &OtherFile{}
	fmt.Printf("%T\n", f.Close())

	// Interface satisfaction
	var w Writer
	w = os.Stdout         // *os.File has method Write()
	w = new(bytes.Buffer) // *bytes.Buffer has method Write()
	//w = time.Hour         // !!! compile time error: no method Write()
	w.Write([]byte("abcd"))

	var f2 File
	f2 = os.Stdout
	w = f2
	//f2 = w  // !!! compile time error: w has no methods Read() and Close()

	// Maps of interfaces
	m := map[*Writer]struct{ x, y float64 } {&w: {5,3}}
	var _ map[string]interface{}
	fmt.Printf("%+v\n", m)

	// Interfaces embedding
	fb := Football{Ball{Radius: 5, Material: "leather"}}
	fb.Bounce()
}
