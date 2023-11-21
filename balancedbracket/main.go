package main

import (
	"fmt"
)

func isBalanced(input string) string {
	stack := []rune{}
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range input {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	//Testcases
	inputs := []string{
		"{ [ ( ) ] }",
		"{ [ ( ] ) }",
		"{ ( ( [ ] ) [ ] ) [ ] }",
	}

	for _, input := range inputs {
		result := isBalanced(input)
		fmt.Printf("Input: %s \n Output: %s\n\n", input, result)
	}
}
