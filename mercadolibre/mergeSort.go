package main

import "fmt"

func mergeArrays(a []int32, b []int32) []int32 {
	r := []int32{}
	al := len(a)
	bl := len(b)

	for i, j := 0, 0; i < al && j < bl; {
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
		} else if a[i] > b[j] {
			r = append(r, b[j])
			j++
		} else if a[i] == b[j] {
			r = append(r, a[i], b[j])
			i++
			j++
		}

		if i == al {
			r = append(r, b[j:]...)
			return r
		} else if j == bl {
			r = append(r, a[i:]...)
			return r
		}
	}
	return r
}

func main() {
	a := []int32{1, 5, 7, 7}
	b := []int32{0, 1, 2, 3}
	fmt.Println(mergeArrays(a, b))
}
