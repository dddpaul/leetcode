package main

import (
	"fmt"
)

/**
 * https://leetcode.com/explore/featured/card/june-leetcoding-challenge-2021/603/week-1-june-1st-june-7th/3769/
 */

type Solver interface {
	LongestConsecutive(nums []int) int
}

type DirectSolver struct{}

func (ds DirectSolver) LongestConsecutive(nums []int) int {
	m := make(map[int]int)
	for _, i := range nums {
		vPrev, okPrev := m[i-1]
		vNext, okNext := m[i+1]
		count := vPrev + vNext + 1

		m[i] = count
		if okPrev {
			m[i-1] = count
		}
		if okNext {
			m[i+1] = count
		}
	}

	fmt.Printf("m = %v\n", m)

	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	return max
}
