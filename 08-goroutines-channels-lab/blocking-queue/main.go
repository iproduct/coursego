package main

import (
	"math/rand"
	"sync"
)

type BlockingQueue struct {
	cond     *sync.Cond
	elements []int64
}

func NewBlockingQueue() *BlockingQueue {
	return &BlockingQueue{cond: sync.NewCond(&sync.Mutex{})}
}

func (q *BlockingQueue) Push(element int64) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.elements = append(q.elements, element)
	q.cond.Signal()
}

func (q *BlockingQueue) Pop() int64 {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for len(q.elements) == 0 {
		q.cond.Wait()
	}

	var head int64
	head, q.elements = q.elements[0], q.elements[1:]
	return head
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	blockingQueue := NewBlockingQueue()
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				blockingQueue.Push(rand.Int63n(100))
			}
		}()
	}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				_ = blockingQueue.Pop()
				//fmt.Printf("%v, ", val)
			}
		}()
	}

	wg.Wait()
}
