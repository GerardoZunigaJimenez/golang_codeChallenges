package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 4)
	finish := make(chan bool)
	array := []int{1, 2, 3, 4}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, value := range array {
			ch <- &value
		}
		finish <- true
		return
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case value := <-ch:
				fmt.Println(*value)
			case <-finish:
				fmt.Println("finish")
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println("ending")
}
