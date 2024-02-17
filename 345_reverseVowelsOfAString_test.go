package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	assert.Equal(t, "holle", reverseVowels("hello"))
	assert.Equal(t, "leotcede", reverseVowels("leetcode"))

	assert.Equal(t, "oteta", reverseVowels("ateto"))
}

func reverseVowels(s string) string {
	lo, hi := 0, len(s)-1
	runes := make([]rune, len(s))
	for i, r := range s {
		runes[i] = r
	}

	for lo < hi {
		if !isVowel(runes[hi]) {
			hi--
			continue
		}

		if !isVowel(runes[lo]) {
			lo++
			continue
		}

		runes[lo], runes[hi] = runes[hi], runes[lo]
		lo++
		hi--
	}

	return string(runes)
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
