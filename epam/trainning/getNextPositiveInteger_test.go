package trainning

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	cases :=
		[]struct {
			description string
			input       []int
			expected    int
		}{
			{
				description: "inner value",
				input:       []int{1, 2, 3, 4, 6, 1},
				expected:    5,
			},
			{
				description: "next latest value",
				input:       []int{2, 3},
				expected:    4,
			},
			{
				description: "negatives",
				input:       []int{-1, -3},
				expected:    1,
			},
		}

	t.Parallel()
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			assert.Equal(t, c.expected, Solution(c.input))
		})
	}
}
