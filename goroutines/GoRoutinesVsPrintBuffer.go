package main

import (
	"fmt"
	"time"
)

const MaxGoroutines = 3000000

func main() {
	start := time.Now()

	c := make(chan int)
	i := 0
	for i < MaxGoroutines {
		go func(i int, c chan int) {
			c <- i + 1
		}(i, c)
		i++
	}

	i = 0
	for i < MaxGoroutines {
		_ = <-c
		i++

	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
