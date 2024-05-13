package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	// leetcode test cases
	assert.Equal(t, true, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	assert.Equal(t, false, canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))

	// corner cases
	assert.Equal(t, false, canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	assert.Equal(t, true, canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	assert.Equal(t, true, canPlaceFlowers([]int{1, 0, 0, 0, 0}, 1))
}
