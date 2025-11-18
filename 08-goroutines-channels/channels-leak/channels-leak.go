package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func firstLeaked(dbs []string) string {
	result := make(chan string)
	for _, db := range dbs {
		go func(db string) {
			result <- request(db)
		}(db)
	}
	return <-result
}

func firstFixed(dbs []string) string {
	result := make(chan string, 3)
	for _, db := range dbs {
		go func(db string) {
			result <- request(db)
		}(db)
	}
	return <-result
}

func firstFixed2(dbs []string) string {
	result := make(chan string)
	for _, db := range dbs {
		go func(db string) {
			result <- request(db)
		}(db)
	}

	defer func() {
		for i := 0; i < len(dbs)-1; i++ {
			<-result
		}
	}()
	return <-result
}

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println(runtime.NumGoroutine())
	res := firstFixed2([]string{"db1", "db2", "db3"})
	fmt.Println("res", res)

	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func request(db string) string {
	time.Sleep(time.Duration(rand.Intn(1000) * int(time.Millisecond)))
	return db
}
