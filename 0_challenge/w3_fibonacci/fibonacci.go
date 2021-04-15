package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3709/
 */

type Solver interface {
	Fib(n int) int
}

type RecursiveSolver struct{}

func (s RecursiveSolver) Fib(n int) int {
	var fib func(int) int
	fib = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	return fib(n)
}

type LinearSolver struct{}

func (s LinearSolver) Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	fib0, fib1 := 0, 1
	for i := 1; i < n; i = i + 1 {
		fib0, fib1 = fib1, fib0+fib1
	}
	return fib1
}

type RecursiveLinearSolver struct{}

func (s RecursiveLinearSolver) Fib(n int) int {
	var fib func(int) int
	fib0, fib1 := 0, 1
	fib = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		fib0, fib1 = fib1, fib0+fib(n-1)
		return fib1
	}
	return fib(n)
}
