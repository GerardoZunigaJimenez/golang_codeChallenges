package main

import (
	"fmt"
	"strings"
)

func verificaCapicua(cadena string) bool {
	cadena = strings.ReplaceAll(cadena, ".", "")
	cadena = strings.ReplaceAll(cadena, "-", "")
	cadena = strings.ReplaceAll(cadena, " ", "")

	l := len(cadena)
	for i, j := 0, l-1; i != j; i++ {
		b := string(rune(cadena[i]))
		e := string(rune(cadena[j]))

		if !strings.EqualFold(b, e) {
			return false
		}
		j--
	}
	return true
}

func main() {
	fmt.Println(verificaCapicua("- La ruta nos aportó otro paso natural. -"))
}
