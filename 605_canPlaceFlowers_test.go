package main

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

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	ln := len(flowerbed)
	for k, v := range flowerbed {
		if v == 0 &&
			(k == 0 || flowerbed[k-1] == 0) &&
			(k == ln-1 || flowerbed[k+1] == 0) {
			flowerbed[k] = 1
			count++
			if count >= n {
				return true
			}
		}
	}

	return count >= n
}
