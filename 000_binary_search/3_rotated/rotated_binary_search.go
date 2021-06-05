package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/125/template-i/952/
 */

type Solver interface {
	Search(nums []int, target int) int
}

type DirectSolver struct{}

/**
 * Return index of greatest number in rotated array
 */
func find_max_index(nums []int) int {
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

	find_max_index := func(nums []int) int {
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

	search := func(nums []int, target int) int {
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

	pivot := find_max_index(nums)
	arrs := [][]int{
		nums[:pivot+1],
		nums[pivot+1:],
	}

	offset := 0
	for _, arr := range arrs {
		i := search(arr, target)
		if i > -1 {
			return offset + i
		}
		offset = offset + len(arr)
	}

	return -1
}
