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
		fmt.Printf("i = %d, step = %d\n", i, step)
		if nums[i] != target && (i <= 0 || i >= len(nums)-1) {
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
	arr := []int{-1, -1}
	return arr
}
