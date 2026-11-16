package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrUserNotFound = errors.New("user not found in database")

type MyError struct {
	Time   time.Time
	Reason string
}

func (myerr *MyError) Error() string {
	return fmt.Sprintf("[%v] %s", myerr.Time, myerr.Reason)
}

func fetchUserFromDB(id int) error {
	// Simulating a database lookup failure
	return &MyError{time.Now(), "user not found in database"}
}

func getUserProfile(id int) error {
	err := fetchUserFromDB(id)
	if err != nil {
		// %w wraps ErrUserNotFound, creating an error chain
		return fmt.Errorf("failed to get profile for ID %d: %w", id, err)
	}
	return nil
}

func main() {
	err := getUserProfile(42)
	//var cause *MyError
	//if errors.As(err, &cause) {
	if err != nil {
		fmt.Printf("[%T]:%v\n", err, err)
	}
}
