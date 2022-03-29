package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed version.txt
var b []byte

func main() {
	s := string(b)
	fmt.Println(s)
	vars := strings.Split(s, "\n")
	for i, s := range vars {
		fmt.Println("build version:", i, " - ", s)
	}
}
