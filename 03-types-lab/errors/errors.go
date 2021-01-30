package main

import (
	"fmt"
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

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

//var deleteTempFiles func()
//filename := "myfile.txt"

//for try := 0; try < 2; try++ {
//	file, err := os.Create(filename)
//	if err == nil {
//		return
//	}
//	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
//		deleteTempFiles()  // Recover some space.
//		continue
//	}
//	fmt.Println(file)
//	return
//}
//}
