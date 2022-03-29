package main

import (
	"fmt"
	"strings"
)

func main() {
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	var keys []string
	for k := range commits {
		keys = append(keys, k)
	}

	s := strings.Join(keys, ",\n\t-->")

	err := fmt.Errorf("Errores:\n%s", s)
	fmt.Println(err.Error())
}