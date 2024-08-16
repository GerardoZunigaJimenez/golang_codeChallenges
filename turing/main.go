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

// TODO it is not working
func main() {
	//input {1, 2, 3}
	//output {1, 2, 3}, {1, 3, 2}, {2, 3, 1},{2, 1, 3}, {3,1,2}, {3,2,1}
	fmt.Println(getCombinations([]int{1, 2, 3}))
}
