package main

import (
	"fmt"
	"testing"
)

func TestMergeStrings(t *testing.T) {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	n := len([]rune(word1))
	m := len([]rune(word2))
	i := 0
	var res []uint8

	for i < n || i < m {
		if i < n {
			res = append(res, word1[i])
		}
		if i < m {
			res = append(res, word2[i])
		}
		i++
	}

	return string(res)
}
