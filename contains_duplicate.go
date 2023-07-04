package main

// ContainsDuplicate
/*
https://leetcode.com/problems/contains-duplicate/submissions/986514156/
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Input: nums = [1,2,3,1]
Output: true

Input: nums = [1,2,3,4]
Output: false
*/
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, i := range nums {
		if seen[i] {
			return true
		} else {
			seen[i] = true
		}
	}

	return false
}
