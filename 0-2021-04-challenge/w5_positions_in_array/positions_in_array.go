package main

import "fmt"

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/597/week-5-april-29th-april-30th/3725/
 */

type Solver interface {
	SearchRange(nums []int, target int) []int
}

type DirectSolver struct{}

func (ds DirectSolver) SearchRange(nums []int, target int) []int {
	arr := []int{-1, -1}

	found1 := false
	for i := 0; i < len(nums); i = i + 1 {
		if nums[i] == target {
			if !found1 {
				arr[0] = i
				arr[1] = i
				found1 = true
			} else {
				arr[1] = i
			}
		}
	}

	return arr
}

type BisectSolver struct{}

func (bs BisectSolver) SearchRange1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	step := len(nums) / 2
	var search func(int) int
	search = func(i int) int {
		step = step / 2
		if step == 0 {
			step = 1
		}
		fmt.Printf("i = %d, step = %d, target = %d\n", i, step, target)
		if i < 0 || i > len(nums)-1 {
			return -1
		}
		if len(nums) > 2 && (i == 0 || i == len(nums)-1) && nums[i] != target {
			return -1
		}
		if nums[i] > target {
			return search(i - step)
		} else if nums[i] < target {
			return search(i + step)
		}
		return i
	}
	return search(step)
}

func (bs BisectSolver) SearchRange(nums []int, target int) []int {
	searchRange1 := func(nums []int, target int) int {
		if len(nums) == 0 {
			return -1
		}
		step := len(nums) / 2
		var search func(int) int
		search = func(i int) int {
			step = step / 2
			if step == 0 {
				step = 1
			}
			fmt.Printf("i = %d, step = %d, target = %d\n", i, step, target)
			if i < 0 || i > len(nums)-1 {
				return -1
			}
			if len(nums) > 2 && (i == 0 || i == len(nums)-1) && nums[i] != target {
				return -1
			}
			if nums[i] > target {
				return search(i - step)
			} else if nums[i] < target {
				return search(i + step)
			}
			return i
		}
		return search(step)
	}

	i := searchRange1(nums, target)
	j := i
	arr := []int{-1, -1}
	if i == -1 {
		return arr
	}
	found1, found2 := false, false
	for {
		if nums[i] == target {
			arr[0] = i
		} else {
			found1 = true
		}
		if nums[j] == target {
			arr[1] = j
		} else {
			found2 = true
		}
		if (found1 && found2) || (i == 0 && j == len(nums)-1) {
			break
		}
		if i > 0 {
			i = i - 1
		}
		if j < len(nums)-1 {
			j = j + 1
		}
	}

	return arr
}
