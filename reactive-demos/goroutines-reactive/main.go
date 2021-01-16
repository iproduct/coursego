package main

import (
	"fmt"
	"github.com/iproduct/coursego/reactive-demos/goroutines"
	"github.com/iproduct/coursego/reactive-demos/iot"
	"math/rand"
	"sync"
	"time"
)

func main() {
	mergedEvents := goroutines.FanIn(goroutines.ProduceEvents(iot.Distance, 10), goroutines.ProducePings(1, 10))
	workerChannels := goroutines.FanOut(goroutines.AccumulateDistance(goroutines.FilterDistance(mergedEvents)), 5)
	var wg sync.WaitGroup
	wg.Add(len(workerChannels))
	for index, wChan := range workerChannels {
		go func(index int, wChan chan iot.IotEvent) {
			for event := range wChan {
				time.Sleep(time.Duration(rand.Intn(1000) + 200) * time.Millisecond) // simulate long processing task
				fmt.Printf("Worker %v successfully processed event: %v\n", index, event)
			}
			wg.Done() // signal the work group that processing has finished
		}(index, wChan)
	}
	//for event := range AccumulateDistance(FilterDistance(mergedEvents)) {
	//for event := range mergedEvents {
	//	fmt.Println(event) // accumulated distances
	//}
	wg.Wait()
}
