package main

/**
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/
 */

type Solver interface {
	LengthOfLongestSubstring(s string) int
}

type DirectSolver struct{}

func (ds DirectSolver) LengthOfLongestSubstring(s string) int {
	return 0
}
