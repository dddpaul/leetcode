package main

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
