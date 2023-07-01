package main

// PartitionString
/*
* https://leetcode.com/problems/optimal-partition-of-string/description/
* Given a string s, partition the string into one or more substrings such that the characters in each substring are unique. That is, no letter appears in a single substring more than once.
* Return the minimum number of substrings in such a partition.
* Note that each character should belong to exactly one substring in a partition.
*
*Input: s = "abacaba"
*Output: 4
*Explanation:
*Two possible partitions are ("a","ba","cab","a") and ("ab","a","ca","ba").
*It can be shown that 4 is the minimum number of substrings needed.
 */
func PartitionString(s string) int {
	index := make(map[rune]bool)
	r := 1
	for _, c := range s {
		if index[c] == false {
			index[c] = true
		} else {
			r++
			index = make(map[rune]bool)
			index[c] = true
		}
	}

	return r
}
