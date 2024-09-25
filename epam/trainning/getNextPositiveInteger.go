package trainning

import (
	"sort"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func Solution(A []int) int {
	sort.Ints(A)
	l := len(A)
	var r int

	for i := 1; i < l; i++ {
		d := (A[i] - A[i-1])
		if d > 1 {
			r = A[i-1] + 1
		}
		if i == l && A[i] < 0 {
			r = 1
		}
	}

	if r < 0 {
		return 1
	} else if r == 0 {
		return A[l-1] + 1
	}
	return r
}
