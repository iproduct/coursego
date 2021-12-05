package main

import (
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	return make(map[string]string)
}

func requests() chan int {
	return make(chan int)
}

func main() {
	var config atomic.Value // holds current server configuration
	// Create initial config value and store into config.
	config.Store(loadConfig())
	for i := 0; i < 10; i++ {
		go func() {
			// Reload config every 10 seconds
			// and update config value with the new version.
			for {
				time.Sleep(10 * time.Millisecond)
				config.Store(loadConfig())
			}
		}()
	}
	// Create worker goroutines that handle incoming requests
	// using the latest config value.
	for i := 0; i < 100; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				// Handle request r using config c.
				_, _ = r, c
			}
		}()
	}
}
