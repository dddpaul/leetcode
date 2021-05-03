package main

import (
	"fmt"
	"math"
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
	step := len(nums) / 2
	count := 0
	limit := int(math.Ceil(math.Log2(float64(len(nums))))) + 1
	var search func(int) int
	search = func(i int) int {
		step = step / 2
		if step == 0 {
			step = 1
		}
		fmt.Printf("count = %d, i = %d, step = %d, target = %d\n", count, i, step, target)
		if (i < 0 || i > len(nums)-1) || count > limit {
			return -1
		}
		count = count + 1
		if nums[i] > target {
			return search(i - step)
		} else if nums[i] < target {
			return search(i + step)
		}
		return i
	}
	return search(step)
}
