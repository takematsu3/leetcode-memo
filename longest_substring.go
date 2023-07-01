package main

// lengthOfLongestSubstring
/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
Given a string s , find the length of the longest substring without repeating characters.

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	for i, ci := range s {
		count := 1
		seen := make(map[rune]bool)
		seen[ci] = true
		for _, cj := range s[i+1:] {
			if _, ok := seen[cj]; ok {
				break
			}
			seen[cj] = true
			count++
			if max < count {
				max = count
			}
		}
	}

	return max
}
