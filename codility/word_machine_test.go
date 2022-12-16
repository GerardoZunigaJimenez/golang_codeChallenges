package solution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Solution(t *testing.T) {
	assert.Equal(t, 7, Solution("4 5 6 - 7"))
	assert.Equal(t, 7, Solution("13 DUP 4 POP 5 DUP + DUP + -"))
	assert.Equal(t, -1, Solution("5 6 + -"))
	assert.Equal(t, -1, Solution("3 DUP 5 - -"))
	assert.Equal(t, -1, Solution("1048575 DUP +"))
	assert.Equal(t, -1, Solution("2 1 -"))
}
