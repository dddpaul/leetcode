package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/125/template-i/950/
 */

type Solver interface {
	MySqrt(x int) int
}

type DirectSolver struct{}

func (ds DirectSolver) MySqrt(x int) int {
	l, r := 0, x
	for {
		i := l + (r-l)/2
		if l > r {
			return i - 1
		}
		if i*i > x {
			r = i - 1
		} else if i*i < x {
			l = i + 1
		} else {
			return i
		}
	}
}
