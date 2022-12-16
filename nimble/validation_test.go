package nimble

import (
	"fmt"
	"testing"
)

func Test_validateNumberPower(t *testing.T) {
	fmt.Println(123456, validateNumberPower(123456))
	fmt.Println(998001, validateNumberPower(998001))
}
