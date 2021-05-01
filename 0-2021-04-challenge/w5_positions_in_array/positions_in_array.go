package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/597/week-5-april-29th-april-30th/3725/
 */

type Solver interface {
	SearchRange(nums []int, target int) []int
}

type DirectSolver struct{}

func (ds DirectSolver) SearchRange(nums []int, target int) []int {
	return nil
}
