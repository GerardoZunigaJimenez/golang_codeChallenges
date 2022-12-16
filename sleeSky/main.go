// input:
// [
//  [1,2,3],
//  [4,5,6]
// ]

// output: cols are rows
// [
//   [1,4],
//   [2,5],
//   [3,6]
// ]
package main

import (
	"fmt"
)

func main() {
	s := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	fmt.Println(process(s))
}

func process(input [][]int) (output [][]int) {
	l := len(input[0])
	output = make([][]int, l)

	for i := 0; i < len(input); i++ {
		ts := input[i]
		ll := len(ts)
		for j := 0; j < ll; j++ {
			output[j] = append(output[j], ts[j])
		}
	}

	return
}
