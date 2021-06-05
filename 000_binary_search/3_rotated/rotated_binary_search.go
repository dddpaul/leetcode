package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/125/template-i/952/
 */

type Solver interface {
	Search(nums []int, target int) int
}

type DirectSolver struct{}

func find_pivot(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	l, r := 0, len(nums)-1
	for {
		i := l + (r-l)/2
		if i < len(nums)-1 && nums[i] > nums[i+1] {
			return i
		}
		if nums[0] > nums[i] {
			r = i - 1
		} else {
			l = i + 1
		}
		if l > r {
			return -1
		}
	}
}

func (ds DirectSolver) Search(nums []int, target int) int {
	return 0
}
