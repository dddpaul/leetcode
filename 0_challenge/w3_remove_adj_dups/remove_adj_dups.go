package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3710/
 */

type Solver interface {
	RemoveDuplicates(s string, k int) string
}

type DirectSolver struct{}

func (ds DirectSolver) RemoveDuplicates(s string, k int) string {
	return ""
}
