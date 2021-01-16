package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"strconv"
	"strings"
	"time"
)

func main() {
	observable := rxgo.Just("Hello", "Reactive", "World", "from", "RxGo", "!", errors.New("Foo Error"))() //.
	//	Map(ToUpper). // map to upper case
	//	Filter(LengthGreaterThan4) // greaterThan4 func filters values > 4

	interval := rxgo.Interval(rxgo.WithDuration(500 * time.Millisecond))

	result := interval.ZipFromIterable(observable, zipper) //,  rxgo.WithPublishStrategy())


	ch := result.Observe()
	go func() {
		for item := range ch {
			//item := <-ch
			if item.Error() {
				fmt.Println("Error:", item.E.Error())
			} else {
				fmt.Println(item.V)
			}
		}
	}()
	disposed := result.ForEach(func(v interface{}) {
		fmt.Printf("received: %v\n", v)
	}, func(err error) {
		fmt.Printf("error: %e\n", err)
	}, func() {
		fmt.Println("observable is closed")
	})
	//result.Connect(context.Background())
	<-disposed // wait until finish
}

// helper functions
func ToUpper(ctx context.Context, item interface{}) (interface{}, error) {
	return strings.ToUpper(item.(string)), nil
}

func LengthGreaterThan4(item interface{}) bool {
	return len(item.(string)) > 4
}

func zipper(_ context.Context, i1 interface{}, i2 interface{}) (interface{}, error) {
return strconv.Itoa(i1.(int)) + ": " +  i2.(string), nil
}