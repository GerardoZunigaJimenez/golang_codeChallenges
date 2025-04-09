package recursiveFibonacci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecursiveFibonacci(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		inputLimit int
		validation func([]int)
	}{
		{
			name:       "Success Flow 5 max",
			inputLimit: 5,
			validation: func(result []int) {
				assert.Equal(t, []int{0, 1, 1, 2, 3, 5}, result)
			},
		},
		{
			name:       "Success Flow 5 max",
			inputLimit: 6,
			validation: func(result []int) {
				assert.Equal(t, []int{0, 1, 1, 2, 3, 5}, result)
			},
		},
		{
			name:       "Success Flow 5 max",
			inputLimit: 8,
			validation: func(result []int) {
				assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8}, result)
			},
		},
		{
			name:       "Success Flow 13 max",
			inputLimit: 17,
			validation: func(result []int) {
				assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8, 13}, result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Fibonacci(tt.inputLimit)
			tt.validation(r)
		})
	}
}

func TestRecursiveFibonacciMax(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		inputLimit int
		validation func([]int)
	}{
		{
			name:       "Success Flow 5 max",
			inputLimit: 5,
			validation: func(result []int) {
				assert.Equal(t, []int{0, 1, 1, 2, 3, 5}, result)
			},
		},
		{
			name:       "Success Flow 5 max",
			inputLimit: 6,
			validation: func(result []int) {
				assert.Equal(t, []int{0, 1, 1, 2, 3, 5}, result)
			},
		},
		{
			name:       "Success Flow 5 max",
			inputLimit: 8,
			validation: func(result []int) {
				assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8}, result)
			},
		},
		{
			name:       "Success Flow 13 max",
			inputLimit: 17,
			validation: func(result []int) {
				assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8, 13}, result)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := recursiveFibonacciMax(nil, tt.inputLimit)
			tt.validation(r)
		})
	}
}
