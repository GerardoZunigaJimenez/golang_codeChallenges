package main

import (
	_ "embed"
	"fmt"
)

//go:generate bash get_version.sh
//go:embed version.txt
var version string

func main() {
	fmt.Println("build version: ", version)
}
