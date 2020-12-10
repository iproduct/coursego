package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
)

type Customer struct {
	ID             int
	Name, LastName string
	Age            int
	TaxNumber      string
}

func main() {
	var wg sync.WaitGroup

	// Create the input channel
	ch := make(chan rxgo.Item)

	// Data producer: go producer(ch)
	wg.Add(1)
	go producer(ch, &wg)

	// Create an Observable
	observable := rxgo.FromChannel(ch)

	connectable := observable.
		Filter(func(item interface{}) bool {
			// Filter operation
			customer := item.(Customer)
			return customer.Age > 15
		}).
		Map(func(_ context.Context, item interface{}) (interface{}, error) {
				// Enrich operation
				customer := item.(Customer)
				taxNumber, err := getTaxNumber(customer)
				if err != nil {
					return nil, err
				}
				customer.TaxNumber = taxNumber
				return customer, nil
			},
			// Create multiple instances of the map operator
			rxgo.WithPool(4),
			////Serialize the items emitted by their Customer.ID
			//rxgo.Serialize(func(item interface{}) int {
			//	customer := item.(Customer)
			//	return customer.ID
			//}),
			rxgo.WithBufferedChannel(1),
			rxgo. WithPublishStrategy(),
			rxgo.WithBackPressureStrategy(rxgo.Block))

	wg.Add(1)
	go func() {
		defer wg.Done()
		for item := range connectable.Observe() {
			if item.Error() {
				fmt.Println("Error:", item.E.Error())
			} else {
				fmt.Println(item.V)
			}
		}
	}()

	complete := connectable.ForEach(func(v interface{}) {
		fmt.Printf("received: %v\n", v)
	}, func(err error) {
		fmt.Printf("error: %e\n", err)
	}, func() {
		fmt.Println("observable is closed")
	})
	connectable.Connect(context.Background())
	<-complete

	wg.Wait()
}

func producer(ch chan<- rxgo.Item, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		customer1 := Customer{10 + i, "Georgi", "Hristov", 20 + i, "123456789"}
		customer2 := Customer{10 - i, "Georgi", "Hristov", 20 - i, "123456789"}
		ch <- rxgo.Of(customer1)
		ch <- rxgo.Of(customer2)
	}
	close(ch)
}

func getTaxNumber(customer Customer) (string, error) {
	return customer.TaxNumber, nil
}
