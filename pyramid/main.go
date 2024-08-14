package main

import "fmt"

func func1(input []int) (output []int) {
	counter := 0
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			input = append(input[0:i], input[i+1:]...)
			counter++
			i++
		}

	}
	for i := 0; i < counter; i++ {
		input = append(input, 0)
	}

	return input
}

func main() {
	fmt.Println(func1([]int{1, 2, 0, 4, 3, 0, 5, 0}))
}
