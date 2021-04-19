package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3713/
 */

type Solver interface {
	CombinationSum4(nums []int, target int) int
}

type DirectSolver struct{}

func (ds DirectSolver) CombinationSum4(nums []int, target int) int {
	return 0
}
