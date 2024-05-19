package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroesCase1(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	assert.Equal(
		t,
		[]int{1, 3, 12, 0, 0},
		a,
	)
}

func TestMoveZeroesCase2(t *testing.T) {
	a := []int{0}
	moveZeroes(a)
	assert.Equal(
		t,
		[]int{0},
		a,
	)
}

func TestMoveZeroesCase3(t *testing.T) {
	a := []int{0, 1, 0, 3, 12, 0, 0, 0, 0, 150}
	moveZeroes(a)
	assert.Equal(
		t,
		[]int{1, 3, 12, 150, 0, 0, 0, 0, 0, 0},
		a,
	)
}

func TestMoveZeroesCase4(t *testing.T) {
	a := []int{1, 3, 12, 150}
	moveZeroes(a)
	assert.Equal(
		t,
		[]int{1, 3, 12, 150},
		a,
	)
}
