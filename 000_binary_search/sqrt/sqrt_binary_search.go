package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/125/template-i/950/
 */

type Solver interface {
	MySqrt(x int) int
}

type DirectSolver struct{}

func (ds DirectSolver) MySqrt(x int) int {
	return 0
}
