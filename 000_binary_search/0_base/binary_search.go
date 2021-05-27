package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/138/background/1038/
 */

type Solver interface {
	Search(nums []int, target int) int
}

type RecursiveSolver struct{}

func (rs RecursiveSolver) Search(nums []int, target int) int {
	if len(nums) == 0 || (nums[0] > target && nums[len(nums)-1] < target) {
		return -1
	}
	i := 0
	var search func(int, int) int
	search = func(l int, r int) int {
		i = l + (r-l)/2
		if l > r {
			return -1
		}
		if nums[i] > target {
			return search(l, i-1)
		} else if nums[i] < target {
			return search(i+1, r)
		}
		return i
	}
	return search(0, len(nums)-1)
}

type DirectSolver struct{}

func (ds DirectSolver) Search(nums []int, target int) int {
	if len(nums) == 0 || (nums[0] > target && nums[len(nums)-1] < target) {
		return -1
	}
	l, r := 0, len(nums)-1
	for {
		i := l + (r-l)/2
		if nums[i] > target {
			r = i - 1
		} else if nums[i] < target {
			l = i + 1
		} else {
			return i
		}
		if l > r {
			return -1
		}
	}
}
