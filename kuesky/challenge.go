package main

import (
	"strings"
)

func SearchWordInMatrix(soap []string, word string) (coords []int) {
	//horizontal
	r := searchByMatrix(soap, word)
	if r != nil && r[0] != 0 && r[1] != 0 {
		return r
	}

	// vertical
	soap = transposeVertical(soap)
	r = searchByMatrix(soap, word)
	if r != nil && r[0] != 0 && r[1] != 0 {
		return r
	}

	// inclinado

	return []int{}
}

func searchByMatrix(soap []string, word string) (coords []int) {
	for x, v := range soap {
		if strings.Contains(v, word) {
			ss := strings.Split(v, word)
			bp := len(ss[0])
			return []int{x, bp}
		}
	}
	return []int{}
}

func transposeVertical(slice []string) []string {
	xl := len(slice[0])
	result := make([]string, xl)

	for _, word := range slice {
		for j := 0; j < len(word); j++ {
			r := result[j]
			r += string(rune(word[j]))
			result[j] = r
		}
	}

	return result
}
