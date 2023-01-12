package lyft

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minSubstring(t *testing.T) {
	//assert.Equal(t, "", minSubstring("abcx", "abc"))
	//assert.Equal(t, "", minSubstring("abcx", "abcsdf"))
	//assert.Equal(t, "abc", minSubstring("abc", "abc"))
	//assert.Equal(t, "abbc", minSubstring("abc", "aabbcc"))
	//assert.Equal(t, "abbc", minSubstring("abc", "aaaaaabbccccc"))
	//TODO: This is the only case that is not passing
	assert.Equal(t, "abbcca", minSubstring("aabbc", "abbcca"))
	//assert.Equal(t, "abbc", minSubstring("bca", "aabbcc"))
}
