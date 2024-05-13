package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	assert.Equal(t, "ABC", gcdOfStrings("ABCABC", "ABC"))
	assert.Equal(t, "AB", gcdOfStrings("ABABAB", "ABAB"))
	assert.Equal(t, "", gcdOfStrings("LEET", "CODE"))
}
