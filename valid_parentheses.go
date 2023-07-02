package main

// IsValidParentheses
/*
https://leetcode.com/problems/valid-parentheses/description/
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Input: s = "()"
Output: true

Input: s = "()[]{}"
Output: true

Input: s = "(]"
Output: false
*/
func IsValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []rune
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			l := len(stack)
			if l < 1 {
				return false
			}
			if m[c] != stack[l-1] {
				return false
			}
			stack = stack[:l-1]
		}
	}

	return len(stack) == 0
}
