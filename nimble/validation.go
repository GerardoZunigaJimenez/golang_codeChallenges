package nimble

import "math"

//(For all 6 digit numbers ABCDEF, find the numbers that satisfy the following equation:
//(ABC+DEF)^2=ABCDEF

//123456 -> (123+456)^2 = 579^2= 335241

//998001 = 998 + 001 = 999 * 999 = 998001

func validateNumberPower(input int) bool {
	rq := math.Sqrt(float64(input))
	if math.Mod(rq, 1.0) != 0 {
		return false
	}

	f := int(rq / 1000)
	s := input - (int(rq) * 1000)
	sum := f + s
	r := math.Pow(float64(sum), 2)
	if r == float64(input) {
		return true
	}
	return false
}
