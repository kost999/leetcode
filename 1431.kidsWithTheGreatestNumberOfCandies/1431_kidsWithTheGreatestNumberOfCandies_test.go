package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	assert.Equal(t, []bool{true, true, true, false, true}, kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	assert.Equal(t, []bool{true, false, false, false, false}, kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	assert.Equal(t, []bool{true, false, true}, kidsWithCandies([]int{12, 1, 12}, 10))
}
