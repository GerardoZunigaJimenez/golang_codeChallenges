package main

import (
	"fmt"
	"math/rand"
	"time"
)

// URL: https://dev.to/b0r/go-concurrency-patterns-fan-out-semaphore-1ojf
func main() {
	emps := 10

	// buffered channel, one slot for every goroutine
	// send side can complete without receive (non-blocking)
	ch := make(chan string, emps)

	// max number of RUNNING goroutines at any given time
	// g := runtime.NumCPU()
	g := 2
	// buffered channel, based on the max number of the goroutines in RUNNING state
	// added to CONTROL the number of goroutines in RUNNING state
	sem := make(chan bool, g)

	for e := 0; e < emps; e++ {
		// create 10 goroutines in the RUNNABLE state
		// one for each employee
		go func(emp int) {

			// when goroutine moves from RUNNABLE to RUNNING state
			// send signal/value inside a `sem` channel
			// if `sem` channel buffer is full, this will block
			// e.g. employee takes a seat
			sem <- true
			{
				// simulate the idea of unknown latency (do not use in production)
				// e.g. employee does some work
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

				// once work is done, signal on ch channel
				ch <- fmt.Sprintf("paper %d", emp)
				fmt.Println("employee : sent signal : ", emp)
			}

			// once all work is done pull the value from the `sem` channel
			// give place to another goroutine to do the work
			// e.g. employee stands up and free up seat for another employee
			<-sem
		}(e)
	}

	// wait for all employee work to be done
	for emps > 0 {
		// receive signal sent from the employee
		p := <-ch

		emps--
		fmt.Println(p)
		fmt.Println("manager : received signal : ", emps)

	}
}
