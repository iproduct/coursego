package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"strconv"
	"strings"
)

func main() {
	observable := rxgo.Just("Hello", "Reactive", "World", "from", "RxGo")().
		Map(ToUpper). // map to upper case
		Filter(LengthGreaterThan4) // greaterThan4 func filters values > 4
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
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