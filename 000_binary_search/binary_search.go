package main

import (
	"fmt"
)

/**
 * https://leetcode.com/explore/learn/card/binary-search/138/background/1038/
 */

type Solver interface {
	Search(nums []int, target int) int
}

type DirectSolver struct{}

func (ds DirectSolver) Search(nums []int, target int) int {
	if len(nums) == 0 || (nums[0] > target && nums[len(nums)-1] < target) {
		return -1
	}
	i := 0
	var search func(int, int) int
	search = func(l int, r int) int {
		i = l + (r-l)/2
		fmt.Printf("i = %d, l = %d, r = %d\n", i, l, r)
		if nums[i] == target {
			return i
		}
		if l >= r {
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
