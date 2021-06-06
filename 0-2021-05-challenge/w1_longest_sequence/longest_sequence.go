package main

/**
 * https://leetcode.com/explore/featured/card/june-leetcoding-challenge-2021/603/week-1-june-1st-june-7th/3769/
 */

type Solver interface {
	LongestConsecutive(nums []int) int
}

type DirectSolver struct{}

func (ds DirectSolver) LongestConsecutive(nums []int) int {
	return 0
}
