package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	var d net.Dialer
	ctx, _ := context.WithTimeout(context.Background(), time.Hour)
	conn, err := d.DialContext(ctx, "tcp", "127.0.0.1:8080")
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := conn.Write([]byte("Hello World\n")); err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(conn)
	i := 0
	for i < 2 && scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Message received: ", message)
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning server response: %v\n", err)
	}
	log.Printf("Closing client connection %v\n", conn)
}
