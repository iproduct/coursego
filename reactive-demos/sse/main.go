package sse

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Golang SSE server

// Server address and port
const Address = "localhost:8080"

// ResourcesPath ia basic path to the project in filesystem
const ResourcesPath = "D:/CourseGO/git/coursego/reactive-demos/static"

type Broker struct {

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// Broker stopping channel
	Done chan struct{}

	// stopping the server
	serverDone chan struct{}

	// New client connections
	newClients chan chan []byte

	// Closed client connections
	closingClients chan chan []byte

	// Client connections registry
	clients map[chan []byte]bool
}

func NewServer() (broker *Broker) {
	// Instantiate a broker
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		Done:           make(chan struct{}),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}

	// Set it running - listening and broadcasting events
	go broker.listen()

	return
}

func (broker *Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// Make sure that the writer supports flushing
	flusher, ok := rw.(http.Flusher)

	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Transfer-Encoding", "chunked")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)

	// Signal the broker that we have a new connection
	broker.newClients <- messageChan

	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {

	}()

	// Listen to connection close and un-register messageChan
	// notify := rw.(http.CloseNotifier).CloseNotify()
	notify := req.Context().Done()

	go func() {
		<-notify
		broker.closingClients <- messageChan
	}()

	for {
		// Write to the ResponseWriter
		// Server Sent Events compatible
		message, more := <-messageChan
		if !more {
			fmt.Fprintf(rw, "0\n\n\n\n")
			flusher.Flush()
			return
		}
		//fmt.Fprintf(rw, "data: %s\n\n", message)
		fmt.Fprintf(rw, "%d\n\ndata: %s\n\n", len(message), message)
		// Flush the data immediatly instead of buffering it for later.
		flusher.Flush()
	}
	broker.closingClients <- messageChan
}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:

			// A new client has connected.
			// Register their message channel
			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))
		case s := <-broker.closingClients:
			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case <-broker.Done:
			log.Println("No more data. Stoping the server.")
			// close all client channels
			for client, _ := range broker.clients {
				close(client)
				delete(broker.clients, client)
				log.Printf("Removed client. %d registered clients", len(broker.clients))
			}
			return
		case event := <-broker.Notifier:
			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan, _ := range broker.clients {
				clientMessageChan <- event
			}
		}
	}

}

func StartHttpServer(address string, handler http.Handler) *http.Server {
	server := &http.Server{Addr: address}

	http.Handle("/sse", handler)
	http.Handle("/", http.FileServer(http.Dir(ResourcesPath)))

	go func() {
		// always returns error. ErrServerClosed on graceful close
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	log.Printf("HTTP server started on: %v\n", address)

	// returning reference so caller can call Shutdown()
	return server
}

func StopHttpServer(server *http.Server, timeoutSeconds int ) error {
	//ctx, cancelFunc := context.WithCancel(context.Background())
	ctx := context.Background()
	go func() {
		time.Sleep(time.Duration(timeoutSeconds) * time.Second)
		//cancelFunc()
		server.Close()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}
	log.Printf("HTTP server stopped.")
	return nil
}

