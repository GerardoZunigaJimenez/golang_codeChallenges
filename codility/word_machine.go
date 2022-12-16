package solution

import (
	"strconv"
	"strings"
)

func Solution(S string) int {
	var stack []int
	elements := strings.Fields(S)
	for _, v := range elements {
		if number, err := strconv.Atoi(v); err == nil {
			validateMaxNumber(number)
			stack = append(stack, number)
		} else {
			operation := v
			l := len(stack)
			switch operation {
			case "POP":
				stack = stack[0 : l-1]
			case "DUP":
				stack = append(stack, stack[l-1])
			case "+":
				if l < 2 {
					return -1
				}
				r := validateMaxNumber(stack[l-1] + stack[l-2])
				stack = append(stack[0:l-2], r)
			case "-":
				if l < 2 {
					return -1
				}
				r := validateMaxNumber(stack[l-1] - stack[l-2])
				if r < 0 {
					return -1
				}
				stack = append(stack[0:l-2], r)
			}
		}

	}
	l := len(stack)
	return stack[l-1]
}

func validateMaxNumber(input int) int {
	limit := 1048576 - 1
	if input < limit && input >= 0 {
		return input
	}
	return -1
}
