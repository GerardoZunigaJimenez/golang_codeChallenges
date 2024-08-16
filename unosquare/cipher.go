package unosquare

import (
	"fmt"
)

func Cipher(input string, rotation int) (string, error) {
	if rotation < 0 {
		return "", fmt.Errorf("the shift number is a negative number, it is not possible to continue")
	} else if rotation == 0 {
		// There is not rotattion because the value is zero
		return input, nil
	}

	var output string
	for _, letter := range input {
		if s, ok := isValidLetterRune2(letter, rotation); ok {
			output += string(s)
		} else {
			output += string(letter)
		}

	}

	return output, nil
}

func isValidLetterRune(letter rune, rotation int) (rune, bool) {
	original := "abcdefghijklmnopqrstuvwxyz"
	oL := len(original)
	eRV := rune('z') + int32(oL)
	sRV := rune('a') + int32(oL)
	for _, l := range original {
		if letter == l {
			nR := l + int32(rotation)
			if nR >= eRV {
				nR = sRV + nR - eRV
			}
			return nR, true
		}
	}

	return letter, false
}

func isValidLetterRune2(letter rune, rotation int) (rune, bool) {
	base := "abcdefghijklmnopqrstuvwxyz"
	l := len(base)
	for i, v := range base {
		if letter == v {
			p := i + rotation
			if p > l {
				p = (p - l) - 1
			}
			return rune(base[p]), true
		}
	}

	return letter, false
}
