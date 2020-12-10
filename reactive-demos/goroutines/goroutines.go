package goroutines

import (
	"encoding/json"
	"fmt"
	"github.com/iproduct/coursego/reactive-demos/iot"
	"log"
	"math/rand"
	"sync"
	"time"
)

type EventRegistry map[iot.IotEventType]iot.IotEvent

func ProduceEvents(eventType iot.IotEventType, n int) <-chan iot.IotEvent { //wg *sync.WaitGroup
	prevEvents := EventRegistry{
		iot.Distance:    *iot.NewDistanceEvent(108),
		iot.Temperature: *iot.NewTemperatureEvent(20),
		iot.Humidity:    *iot.NewHumidityEvent(20),
		iot.Light:       *iot.NewLightEvent(20),
		iot.Electricity: *iot.NewElectricityEvent(10),
	}
	//defer wg.Done()
	ch := make(chan iot.IotEvent)
	go func() {
		for i := 1; i <= n; i++ {
			var iotEvent iot.IotEvent
			switch eventType {
			case iot.Distance:
				iotEvent = *iot.NewDistanceEvent(prevEvents[iot.Distance].Readings[0] + rand.Intn(20) - 10)
			case iot.Temperature:
				iotEvent = *iot.NewTemperatureEvent(prevEvents[iot.Temperature].Readings[0] + rand.Intn(10) - 5)
			case iot.Humidity:
				iotEvent = *iot.NewHumidityEvent(prevEvents[iot.Humidity].Readings[0] + rand.Intn(6) - 3)
			case iot.Light:
				iotEvent = *iot.NewLightEvent(prevEvents[iot.Light].Readings[0] + rand.Intn(14) - 7)
			case iot.Electricity:
				iotEvent = *iot.NewElectricityEvent(prevEvents[iot.Electricity].Readings[0] + rand.Intn(6) - 3)
			}
			ch <- iotEvent
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}

func ProducePings(intervalSeconds, forPeriodSeconds int) <-chan iot.IotEvent {
	ticker := time.NewTicker(time.Duration(intervalSeconds) * time.Second)
	done := make(chan bool)
	go func() {
		time.Sleep(time.Duration(forPeriodSeconds) * time.Second)
		done <- true
	}()

	events := make(chan iot.IotEvent)
	go func() {
		defer ticker.Stop()
		defer close(events)
		for {
			select {
			case <-done:
				fmt.Println("Done!")
				return
			case <-ticker.C:
				//fmt.Println("Tick at", t)
				events <- *iot.NewPingEvent()
			}
		}
	}()
	return events
}

func FilterDistance(in <-chan iot.IotEvent) <-chan iot.IotEvent {
	out := make(chan iot.IotEvent)
	go func() {
		for event := range in {
			if event.Type == iot.Distance {
				out <- event
			}
		}
		close(out)
	}()
	return out
}

func AccumulateDistance(in <-chan iot.IotEvent) <-chan iot.IotEvent {
	distance := 0
	out := make(chan iot.IotEvent)
	go func() {
		for ev := range in {
			if ev.Type == iot.Distance && len(ev.Readings) == 1 {
				distance += ev.Readings[0]
				out <- *iot.NewEvent(ev.ID, ev.Type, ev.Timestamp, distance)
			}
		}
		close(out)
	}()
	return out
}

func Jsonify(in <-chan iot.IotEvent) <-chan []byte {
	out := make(chan []byte)
	go func() {
		for ev := range in {
			// IotEvent --> JSON
			data, err := json.Marshal(ev)
			if err != nil {
				log.Fatalf("JSON marshaling failed: %s", err)
			}
			out <- data
		}
		close(out)
	}()
	return out
}

func FanIn(inputs ...<-chan iot.IotEvent) <-chan iot.IotEvent {
	var wg sync.WaitGroup
	out := make(chan iot.IotEvent)

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	send := func(ch <-chan iot.IotEvent) {
		for ev := range ch {
			out <- ev
		}
		wg.Done()
	}

	wg.Add(len(inputs))
	for _, ch := range inputs {
		go send(ch)
	}

	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Fan Out pattern distributes messages between n workers (receivers)
// - e.g in a round-robin fashion
func FanOut(input <-chan iot.IotEvent, n int) []chan iot.IotEvent {
	workers := make([]chan iot.IotEvent, 0)
	for i := 0; i < n; i++ {
		workers = append(workers, make(chan iot.IotEvent))
	}

	go func() {
		// Close all n worker channels when the input is closed
		defer func() {
			for _, worker := range workers {
				close(worker)
			}
		}()

		// Distribute input events in round-robin fashion between workers
		for {
			for _, worker := range workers {
				select {
				case event, ok := <-input:
					if !ok {
						return
					}
					worker <- event
				}
			}
		}
	}()

	return workers // return worker channels
}
