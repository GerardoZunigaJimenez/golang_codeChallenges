package main

import "fmt"

func getCombinations(input []int) (output [][]int) {
	if len(input) == 2 {
		s2 := []int{input[1], input[0]}
		return [][]int{input, s2}
	}

	for i := range input {
		v := input[i]
		ss := append(input[:i], input[i+1:]...)
		fmt.Println(input, ss)
		c := getCombinations(ss)
		for _, sss := range c {
			ssss := append([]int{v}, sss...)
			output = append(output, ssss)
		}
	}
	return output
}

func main() {
	fmt.Println(getCombinations([]int{1, 2, 3}))
}
