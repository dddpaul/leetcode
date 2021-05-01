package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/597/week-5-april-29th-april-30th/3726/
 */

type Solver interface {
	PowerfulIntegers(x int, y int, bound int) []int
}

type DirectSolver struct{}

func (ds DirectSolver) PowerfulIntegers(x int, y int, bound int) []int {
	return nil
}
