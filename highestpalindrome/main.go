package main

import (
	"fmt"
)

func highestPalindromeUtil(s []rune, start, end, k int) string {
	if k < 0 {
		return "-1"
	}
	if start >= end {
		return string(s)
	}

	if s[start] != s[end] {
		maxChar := s[start]
		if s[start] < s[end] {
			maxChar = s[end]
		}

		s[start] = maxChar
		s[end] = maxChar
		k--
	}

	return highestPalindromeUtil(s, start+1, end-1, k)
}

func highestPalindrome(s string, k int) string {
	runes := []rune(s)
	return highestPalindromeUtil(runes, 0, len(runes)-1, k)
}

func main() {
	stringInput := "3943"
	k := 1

	result := highestPalindrome(stringInput, k)
	fmt.Println("Output:", result)
}
