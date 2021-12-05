package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"runtime"
	"sync"
	"time"
)

// step 1: specify the number of jobs
var numJobs = 40

// step 2: specify the job and result
type job struct {
	id    int
	input float64
}
type result struct {
	jobId    int
	workerId int
	output   float64
	err      error
}

func main() {
	jobs := make(chan job)
	results := make(chan result)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(runtime.NumCPU())
	go func() {
		defer close(results)
		for i := 0; i < runtime.NumCPU(); i++ {
			wg.Add(1)
			go func(ctx context.Context, workerId int, jobs <-chan job, results chan<- result) {
				defer wg.Done()
				for job := range jobs {
					// step 3: specify the work for the worker
					val := job.input
					select {
					case <-ctx.Done():
						return
					default:
					}
					for k := 0; k < 10000000; k++ {
						val = math.Pow(val, 1/val)
					}
					var r result
					if job.input > 0 {
						r = result{job.id, workerId, val, nil}
					} else {
						r = result{job.id, workerId, 0, fmt.Errorf("invalid argument error: %f", job.input)}
					}
					select {
					case <-ctx.Done():
						return
					case results <- r:
					}
				}
			}(ctx, i, jobs, results)
		}
		wg.Wait()
	}()

	// step 4: send out jobs
	go func() {
		defer close(jobs)
		for i := 0; i < numJobs; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- job{i, math.Abs(float64(i-14) * math.Pi)}:
				fmt.Printf("Sending job %d to jobs queue.\n", i)
			}
		}
	}()

	// step 5: do something with results
	for r := range results {
		if r.err != nil {
			// do something with error
			log.Printf("Error executing Job %d: %v\n", r.jobId, r.err.Error())
			cancel()
		} else {
			fmt.Printf("Job %d executed by worker %d -> Result: %v\n", r.jobId, r.workerId, r.output)
		}
	}
	fmt.Printf("Final number of goroutines: %d\n", runtime.NumGoroutine())
}
