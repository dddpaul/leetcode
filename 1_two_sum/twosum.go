package main

/**
 * https://leetcode.com/problems/two-sum
 */

type Solver interface {
	TwoSum(nums []int, target int) []int
}

type BruteforceSolver struct{}

func (s BruteforceSolver) TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
