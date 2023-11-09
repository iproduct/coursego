package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

type error interface {
	Error() string
}

type PathError struct {
	Op   string // "open", "unlink", etc.
	Path string // The associated file.
	Err  error  // Returned by the system call.
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	//if err := run(); err != nil {
	//	fmt.Println(err)
	//	if myErr, ok := err.(*MyError); ok {
	//		fmt.Printf("When: %v, What: %v\n", myErr.When, myErr.What)
	//	}
	//
	//}
	fileError()
}

var deleteTempFiles func()

func fileError() {
	filename := "myfile.txt"
	var (
		file *os.File
		err  error
	)
	for try := 0; try < 2; try++ {
		file, err = os.Create(filename)
		if err == nil {
			break
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			deleteTempFiles() // Recover some space.
			continue
		} else {
			return
		}
	}
	fmt.Println(file.Name())
}
