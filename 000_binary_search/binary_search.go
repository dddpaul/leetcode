package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/138/background/1038/
 */

type Solver interface {
	Search(nums []int, target int) int
}

type DirectSolver struct{}

func (ds DirectSolver) Search(nums []int, target int) int {
	return 0
}
