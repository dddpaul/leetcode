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

type HashSolver struct{}

func (s HashSolver) TwoSum(nums []int, target int) []int {
	m := make(map[int][]int)
	var arr []int
	var ok bool
	for i, x := range nums {
		if arr, ok = m[x]; !ok {
			arr = make([]int, 0)
		}
		arr = append(arr, i)
		m[x] = arr
	}
	for i, x := range nums {
		y := target - x
		if arr, ok = m[y]; ok {
			for _, j := range arr {
				if i != j {
					return []int{i, j}
				}
			}
		}
	}
	return []int{}
}
