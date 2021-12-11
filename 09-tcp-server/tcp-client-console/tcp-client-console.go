package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"os"
	"strings"
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

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(con)

	ctx2, cancel2 := context.WithCancel(context.Background())

	// Waiting for the server response
	go func() {
		defer cancel2()
		for {
			serverResponse, err := serverReader.ReadString('\n')

			switch err {
			case nil:
				log.Println(strings.TrimSpace(serverResponse))
			case io.EOF:
				log.Println("server closed the connection")
				return
			default:
				log.Printf("server error: %v\n", err)
				return
			}
		}
	}()

	go func() {
		for {
			// Waiting for the client request
			clientRequest, err := clientReader.ReadString('\n')

			switch err {
			case nil:
				clientRequest := strings.TrimSpace(clientRequest)
				if _, err = con.Write([]byte(clientRequest + "\n")); err != nil {
					log.Printf("failed to send the client request: %v\n", err)
				}
			case io.EOF:
				log.Println("client closed the connection")
				cancel2()
				return
			default:
				log.Printf("client error: %v\n", err)
				cancel2()
				return
			}
		}
	}()
	<-ctx2.Done()
	log.Println("Exiting client ...")
}
