package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3709/
 */

type Solver interface {
	Fib(n int) int
}

type DirectSolver struct{}

func (s DirectSolver) Fib(n int) int {
	var fib func(int) int
	fib = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	return fib(n)
}
