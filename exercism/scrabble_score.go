package main

import (
	"fmt"
	"strings"
)

var scores = map[int][]string{
	1:  []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  []string{"D", "G"},
	3:  []string{"B", "C", "M", "P"},
	4:  []string{"F", "H", "V", "W", "Y"},
	5:  []string{"K"},
	8:  []string{"J", "X"},
	10: []string{"Q", "Z"},
}

func calculateScore(w string) int {
	var score int
	for _, v := range w {
		for sV, sL := range scores {
			for _,l := range sL{
				if strings.EqualFold(string(v), l){
					score += sV
				}
			}
		}
	}

	return score
}

func main() {
	fmt.Println(calculateScore("cabbage"))
}
