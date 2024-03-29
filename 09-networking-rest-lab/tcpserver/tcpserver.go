package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	lConfig := net.ListenConfig{
		Control:   nil,
		KeepAlive: time.Duration(time.Minute),
	}
	ln, err := lConfig.Listen(ctx, "tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Accept connection on port: 8080")

	// accept client connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Client connection error: %v/n", err)
		}
		select {
		case <-ctx.Done():
			log.Println("Server closed.")
			return
		default:
		}
		log.Printf("Handling client connection: %#v\n", conn)
		log.Printf("Current number of goroutines: %d\n", runtime.NumGoroutine())
		go handleConnection(ctx, conn)

	}
}

func handleConnection(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		log.Printf("Message received: %v\n", message)
		if message == ":QUIT" {
			log.Println("Client requested to close the connection so closing.")
			return
		}
		select {
		case <-ctx.Done():
			log.Printf("Connection canceled: %v\n", conn)
			return
		default:
		}
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + ":1\n"))
		conn.Write([]byte(newMessage + ":2\n"))
		if err := scanner.Err(); err != nil {
			log.Printf("Error scanning client request: %v\n", err)
			break
		}
	}
	log.Printf("Closing client connection %v\n", conn)
}
