package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	assert.Equal(t, "holle", reverseVowels("hello"))
	assert.Equal(t, "leotcede", reverseVowels("leetcode"))

	assert.Equal(t, "oteta", reverseVowels("ateto"))
}
