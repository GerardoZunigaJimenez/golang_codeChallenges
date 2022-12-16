package lyft

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minSubstring(t *testing.T) {
	//assert.Equal(t, "", minSubstring("abcx", "abc"))
	//assert.Equal(t, "", minSubstring("abcx", "abcsdf"))
	//assert.Equal(t, "abc", minSubstring("abc", "abc"))
	assert.Equal(t, "abbc", minSubstring("abc", "aabbcc"))
	//assert.Equal(t, "abbc", minSubstring("abc", "aaaaaabbccccc"))
	//assert.Equal(t, "abbcca", minSubstring("aabbc", "abbcca"))
}

func Test_createMapPosition(t *testing.T) {
	m := map[string][]int{"a": {0, 1, 2, 3, 4, 5}, "b": {6, 7}, "c": {8, 9, 10, 11, 12}}
	assert.Equal(t, m, createMapPosition("aaaaaabbccccc"))
	fmt.Println("aaaaaabbccccc", m)
	m = map[string][]int{"a": {0, 3, 6}, "b": {1, 4, 7}, "c": {2, 5, 8}}
	assert.Equal(t, m, createMapPosition("abcabcabc"))
	fmt.Println("abcabcabc", m)
}
