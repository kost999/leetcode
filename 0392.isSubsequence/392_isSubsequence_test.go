package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequenceCase1(t *testing.T) {
	assert.Equal(t,
		true,
		isSubsequence("abc", "ahbgdc"),
	)
}

func TestIsSubsequenceCase2(t *testing.T) {
	assert.Equal(t,
		false,
		isSubsequence("axc", "ahbgdc"),
	)
}

func TestIsSubsequenceCase3(t *testing.T) {
	assert.Equal(t,
		true,
		isSubsequence("", "ahbgdc"),
	)
}

func TestIsSubsequenceCase4(t *testing.T) {
	assert.Equal(t,
		true,
		isSubsequence("b", "abc"),
	)
}
