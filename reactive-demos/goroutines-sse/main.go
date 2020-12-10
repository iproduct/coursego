package main

import (
	"github.com/iproduct/coursego/reactive-demos/goroutines"
	"github.com/iproduct/coursego/reactive-demos/iot"
	"github.com/iproduct/coursego/reactive-demos/sse"
	"log"
)


// Golang SSE server using coroutines

func main() {
	serverDone := make(chan struct{})
	broker := sse.NewServer()

	mergedEvents := goroutines.FanIn(
		goroutines.ProduceEvents(iot.Distance, 100),
		goroutines.ProduceEvents(iot.Temperature,100),
		goroutines.ProduceEvents(iot.Humidity,100),
		goroutines.ProduceEvents(iot.Light,100),
		goroutines.ProduceEvents(iot.Electricity,100),
		goroutines.ProducePings(1, 60))
	jsonEvents := goroutines.Jsonify(mergedEvents)
	go func() {
		for event := range jsonEvents {
			log.Printf("Sending event: %s\n", event)
			broker.Notifier <- event
		}
		broker.Done <- struct{}{}
		serverDone <- struct{}{} // signal the server to stop
	}()
	srv := sse.StartHttpServer(sse.Address, broker)
	// Stop the server when broker is done.
	<-serverDone
	sse.StopHttpServer(srv, 5)
}