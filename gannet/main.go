package main

import "fmt"

func main() {
	fmt.Print("Go ")

	c := make(chan bool)
	go func() {
		fmt.Print("Hello ")
		c <- true
	}()
	<-c
	defer fmt.Print("World")
}
