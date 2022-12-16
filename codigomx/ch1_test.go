package codigomx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckZeros(t *testing.T) {
	input := []int{0, 5, 8, 0, 1, 9, 0}
	expected := []int{0, 0, 0, 5, 8, 1, 9}
	result := CheckZeros(input)

	assert.NotEmpty(t, result)
	assert.Equal(t, expected, result)
}
