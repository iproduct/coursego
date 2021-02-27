package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func panicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}

func timeFunc(f func(map[uint64]string), p map[uint64]string) {
	s := time.Now()
	f(p)
	fmt.Printf("took: %s\n", time.Since(s).String())
}

func populate(m map[uint64]string) {
	sliceCap := uint64(100000000)
	fmt.Printf("filling %d slots\n", sliceCap)
	var i uint64
	for i = 0; i < sliceCap; i++ {
		m[i] = strconv.FormatUint(i, 10)
	}
}

func main() {
	mappo := make(map[uint64]string)
	fmt.Println("populating")
	s := time.Now()
	populate(mappo)
	fmt.Printf("done in %s\n", time.Since(s).String())

	fmt.Println("running prealloc")
	timeFunc(keysPrealloc, mappo)

	fmt.Println("running append")
	timeFunc(keysAppend, mappo)

	fmt.Println("running reflect")
	timeFunc(keysReflect, mappo)
}

func keysPrealloc(m map[uint64]string) {
	k := make([]uint64, len(m))
	var i uint64
	for key := range m {
		k[i] = key
		i++
	}
}

func keysAppend(m map[uint64]string) {
	keys := make([]uint64, 0)
	for key := range m {
		keys = append(keys, key)
	}
}

func keysReflect(m map[uint64]string) {
	reflect.ValueOf(m).MapKeys()
}
