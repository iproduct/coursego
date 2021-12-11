package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Message Received:", message)
		if message == ":QUIT" {
			log.Println("client requested server to close the connection so closing")
			return
		}
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + ":1\n"))
		conn.Write([]byte(newMessage + ":2\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
	log.Printf("Closing client connection: %#v\n", conn)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	lConfig := net.ListenConfig{
		Control:   nil,
		KeepAlive: time.Duration(time.Minute),
	}
	ln, err := lConfig.Listen(ctx, "tcp", "127.0.0.1:8081")
	//ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accept connection on port: 8081")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		select {
		case <-ctx.Done():
			log.Println("Server closed.")
			return
		default:
		}
		fmt.Printf("Calling handleConnection: %#v\n", conn)
		fmt.Printf("Current number of goroutines: %d\n", runtime.NumGoroutine())
		go handleConnection(ctx, conn)
	}
}
