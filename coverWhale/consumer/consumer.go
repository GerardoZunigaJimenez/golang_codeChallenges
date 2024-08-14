package main

import "fmt"

func main() {
	input := []int{5, 8, 9, 6, 4}
	tube := make(chan int)
	finish := make(chan bool)

	go func() {
		for _, v := range input {
			tube <- v
		}
		finish <- true
	}()

	counter := 0
	for {
		select {
		case v := <-tube:
			counter += 2
			fmt.Printf("receiving %d and the new value for counter is %d\n", v, counter)
		case <-finish:
			return
		}
	}

	fmt.Println("finish")
}
