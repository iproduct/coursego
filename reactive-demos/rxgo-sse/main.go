package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iproduct/coursego/reactive-demos/goroutines"
	"github.com/iproduct/coursego/reactive-demos/iot"
	"github.com/iproduct/coursego/reactive-demos/sse"
	"github.com/reactivex/rxgo/v2"
	"log"
)

// Golang SSE server using RxGO

func main() {
	serverDone := make(chan struct{})
	broker := sse.NewServer()

	// Create an Observable
	distance := rxgo.FromChannel(ProduceItems(goroutines.ProduceEvents(iot.Distance, 100)))
	temperature := rxgo.FromChannel(ProduceItems(goroutines.ProduceEvents(iot.Temperature,100)))
	humidity := rxgo.FromChannel(ProduceItems(goroutines.ProduceEvents(iot.Humidity,100)))
	light := rxgo.FromChannel(ProduceItems(goroutines.ProduceEvents(iot.Light,100)))
	electricity := rxgo.FromChannel(ProduceItems(goroutines.ProduceEvents(iot.Electricity,100)))
	pings := rxgo.FromChannel(ProduceItems(goroutines.ProducePings(1, 60)))

	mergedEvents := rxgo.Merge([]rxgo.Observable{distance, temperature, humidity, light, electricity, pings}).
		Filter(func(item interface{}) bool {
			iotEvent := item.(iot.IotEvent)
			return iotEvent.Type == iot.Temperature ||
				iotEvent.Type == iot.Humidity ||
				iotEvent.Type == iot.Light ||
				iotEvent.Type == iot.Electricity ||
				iotEvent.Type == iot.Ping
		}).
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
			event := item.(iot.IotEvent)
			return json.Marshal(event)
		},
			//rxgo.WithPool(4),
			//rxgo.WithBufferedChannel(1),
			//rxgo.WithPublishStrategy(),
			//rxgo.WithBackPressureStrategy(rxgo.Block),
		)

	// Create 500ms interval
	//interval := rxgo.Interval(rxgo.WithDuration(500 * time.Millisecond))
	// Zip produced IotEvent readings with interval => schedule them in 500ms intervals
	//scheduledEvents := interval.ZipFromIterable(observable, zipper)

	go func() {
		for event := range mergedEvents.Observe() {
			if event.Error() {
				fmt.Println("Error:", event.E.Error())
			} else {
				log.Printf("Sending event: %s\n", event)
				broker.Notifier <- event.V.([]byte)
			}
		}
		broker.Done <- struct{}{}
		serverDone <- struct{}{} // signal the server to stop
	}()
	srv := sse.StartHttpServer(sse.Address, broker)
	// Stop the server when broker is done.
	<-serverDone
	sse.StopHttpServer(srv, 5)
}

// Helper functions
func ProduceItems(events <-chan iot.IotEvent) <-chan rxgo.Item {
	out := make(chan rxgo.Item)
	go func() {
		for event := range events {
			out <- rxgo.Of(event)
		}
		close(out)
	}()
	return out
}

func zipper(_ context.Context, i1 interface{}, i2 interface{}) (interface{}, error) {
	return i2, nil
}
