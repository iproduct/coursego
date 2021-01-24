package main

import (
	"fmt"
	"log"
	"time"
)

type MyError struct {
	Time   time.Time
	Reason string
}

func (myerr *MyError) Error() string {
	return fmt.Sprintf("[%v] %s", myerr.Time, myerr.Reason)
}

func badFunction() {
	fmt.Printf("Select Panic type (0=no panic, 1=int, 2=runtime panic)\n")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		panic(MyError{time.Now(), "Invalid choice"})
	case 2:
		var invalid func()
		invalid()
	}
}

func main() {
	defer func() {
		if pan := recover(); pan != nil {
			switch err := pan.(type) {
			default:
				panic(pan)
			case MyError:
				fmt.Printf("Function panicked with an error: %s\n", err)
				fmt.Printf("Timestamp: %v", err.Time)
			}
		}
	}()
	badFunction()
	fmt.Printf("Program exited normally\n")
}

// More examples:
type Work interface{}

var do func(work Work)

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

// Error is the type of a parse error; it satisfies the error interface.
type Error string

func (e Error) Error() string {
	return string(e)
}

//
//// error is a method of *Regexp that reports parsing errors by
//// panicking with an Error.
//func (regexp *Regexp) error(err string) {
//	panic(Error(err))
//}
//
//// Compile returns a parsed representation of the regular expression.
//func Compile(str string) (regexp *Regexp, err error) {
//	regexp = new(Regexp)
//	// doParse will panic if there is a parse error.
//	defer func() {
//		if e := recover(); e != nil {
//			regexp = nil    // Clear return value.
//			err = e.(Error) // Will re-panic if not a parse error.
//		}
//	}()
//	return regexp.doParse(str), nil
//}
