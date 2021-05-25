package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/125/template-i/950/
 */

type Solver interface {
	MySqrt(x int) int
}

type DirectSolver struct{}

func (ds DirectSolver) MySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	l, r := 0, x-1
	for {
		i := l + (r-l)/2
		if i*i > x {
			r = i - 1
		} else if i*i < x {
			l = i + 1
		} else {
			return i
		}
		if l > r {
			return i
		}
	}
}
