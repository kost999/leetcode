package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeStrings(t *testing.T) {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
