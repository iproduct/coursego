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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	con, err := d.DialContext(ctx, "tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer con.Close()

	if _, err := con.Write([]byte("Hello, World!\n")); err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(con)
	i := 0
	for i < 2 && scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Message Received:", message)
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
	//serverReader := bufio.NewReader(con)
	//serverResponse, err := serverReader.ReadString('\n')
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(strings.TrimSpace(serverResponse))
}
