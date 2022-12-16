package codigomx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNotDuplicatedValue(t *testing.T) {
	input := []int{0, 5, 6, 0, 7, 6, 5}
	expected := 7

	r := GetNotDuplicatedValue(input)
	assert.Equal(t, expected, r)
}
