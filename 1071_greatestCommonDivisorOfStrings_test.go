package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	assert.Equal(t, "ABC", gcdOfStrings("ABCABC", "ABC"))
	assert.Equal(t, "AB", gcdOfStrings("ABABAB", "ABAB"))
	assert.Equal(t, "", gcdOfStrings("LEET", "CODE"))
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[:gcd(len([]rune(str1)), len([]rune(str2)))]
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
