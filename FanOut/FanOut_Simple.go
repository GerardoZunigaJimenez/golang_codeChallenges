package main

import (
	"fmt"
	//"math/rand"
	"time"
)

// URL: https://dev.to/b0r/go-orchestration-patterns-fan-out-3d8n
func main() {

	employees := 10

	// make channel of type string which provides signaling semantics
	// buffered channel is used so no goroutine blocks a sending operation
	// if two goroutines send a signal at the same time, channel performs synchronization
	ch := make(chan string)

	for e := 0; e < employees; e++ {

		// start goroutine that does some work for employee
		go func(employee int) {

			// simulate the idea of unknown latency (do not use in production)
			//time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

			ch <- fmt.Sprintf("paper %d", employee)
			fmt.Println("employee : sent signal :", employee)
		}(e)
	}

	// goroutine 'main' => manager
	// goroutines 'main' and employee goroutines are executed in parallel

	// wait for all employee work to be done
	for employees > 0 {
		// receive signal sent from the employee
		p := <-ch

		employees--
		fmt.Println(p)
		fmt.Println("manager : received signal :", p)
	}

	time.Sleep(time.Second)
}
