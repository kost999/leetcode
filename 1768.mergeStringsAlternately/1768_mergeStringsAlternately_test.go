package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeStrings(t *testing.T) {
	assert.Equal(t, "apbqcr", mergeAlternately("abc", "pqr"))
	assert.Equal(t, "apbqrs", mergeAlternately("ab", "pqrs"))
	assert.Equal(t, "apbqcd", mergeAlternately("abcd", "pq"))
}
