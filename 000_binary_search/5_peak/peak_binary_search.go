package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/126/template-ii/948/
 */

type Solver interface {
	FindPeakElement(nums []int) int
}

type DirectSolver struct{}

func (ds DirectSolver) FindPeakElement(nums []int) int {
	return 0
}
