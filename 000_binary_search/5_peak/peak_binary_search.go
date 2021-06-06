package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/126/template-ii/948/
 */

type Solver interface {
	FindPeakElement(nums []int) int
}

type DirectSolver struct{}

func (ds DirectSolver) FindPeakElement(nums []int) int {
	l, r := 0, len(nums)
	for {
		i := l + (r-l)/2
		if (i == 0 || nums[i-1] < nums[i]) && (i == len(nums)-1 || nums[i] > nums[i+1]) {
			return i
		}
		if nums[i-1] > nums[i] {
			r = i
		} else if nums[i] < nums[i+1] {
			l = i + 1
		}
	}
}

type SimplerSolver struct{}

func (ss SimplerSolver) FindPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for {
		i := (l + r) / 2
		if i == len(nums)-1 || nums[i] > nums[i+1] {
			r = i
		} else {
			l = i + 1
		}
		if l == r {
			return l
		}
	}
}
