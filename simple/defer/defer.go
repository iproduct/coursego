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
type MyError2 struct {
	Time   time.Time
	Reason string
}

func (myerr *MyError) Error() string {
	return fmt.Sprintf("[%v] %s", myerr.Time, myerr.Reason)
}
func (myerr *MyError2) Error() string {
	return fmt.Sprintf("[%v] %s", myerr.Time, myerr.Reason)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			e := err.(MyError)
			log.Println("work failed:", e.Error())
		}
	}()
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	panic(MyError{time.Now(), "Invalid choice"})
	fmt.Println("done")
}
