package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3709/
 */

type Solver interface {
	Fib(n int) int
}

type DirectSolver struct{}

func (s DirectSolver) Fib(n int) int {
	return 0
}
