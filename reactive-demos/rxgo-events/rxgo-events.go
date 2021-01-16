package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"time"
	"github.com/iproduct/coursego/reactive-demos/iot"
)

func main() {
	//var wg sync.WaitGroup

	// Create the input channel
	ch := make(chan rxgo.Item)
	// Create an Observable
	observable := rxgo.FromChannel(ch)
	// Create 500ms interval
	interval := rxgo.Interval(rxgo.WithDuration(500 * time.Millisecond))
	// Zip produced IotEvent readings with interval => schedule them in 500ms intervals
	scheduledEvents := interval.ZipFromIterable(observable, zipper)


	// Data produceEvents: go produceEvents(ch)
	//wg.Add(1)
	go produceEvents(ch)



	connectable := scheduledEvents.
		Filter(func(item interface{}) bool {
			// Filter operation
			iotEvent := item.(iot.IotEvent)
			return iotEvent.Type == iot.Distance
		}).
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
				iotEvent := item.(iot.IotEvent)
				return iotEvent, nil
			},
			// Create multiple instances of the map operator
			rxgo.WithPool(4),
			rxgo.WithBufferedChannel(1),
			rxgo.WithPublishStrategy(),
			rxgo.WithBackPressureStrategy(rxgo.Block))

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		connectable.Connect(context.Background())
		for item := range connectable.Observe() {
			if item.Error() {
				fmt.Println("Observer 1: Error:", item.E.Error())
			} else {
				fmt.Println("Observer 1:", item.V)
			}
		}
		fmt.Println("Observer 1: observable is closed")
	}()

	complete := connectable.ForEach(func(v interface{}) {
		fmt.Printf("Observer 2: %v\n", v) // onNext
	}, func(err error) {
		fmt.Printf("Observer 2: Error: %e\n", err) // onError
	}, func() {
		fmt.Println("Observer 2: observable is closed") // onClose
	})
	//ctx, cancelFunc := context.WithCancel(context.Background())
	ctx := context.Background()
	connectable.Connect(ctx)

	// canceling after 5 sec
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	cancelFunc()
	//}()
	<-complete

	//wg.Wait()
}

// Helper functions
func produceEvents(ch chan<- rxgo.Item,) { //wg *sync.WaitGroup
	//defer wg.Done()
	for i := 1; i <= 10; i++ {
		iotEvent := iot.IotEvent{uint64(i), iot.Distance, time.Now(), []int{108 + i} }
		ch <- rxgo.Of(iotEvent)
	}
	close(ch)
}

func zipper (_ context.Context, i1 interface{}, i2 interface{}) (interface{}, error) {
	return i2, nil
}
