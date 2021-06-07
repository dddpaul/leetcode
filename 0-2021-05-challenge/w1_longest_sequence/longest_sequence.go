package main

import (
	"fmt"
	"math"
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
		vPrev, vNext := 0, 0
		okPrev, okNext := false, false

		vPrev, okPrev = m[i-1]
		vNext, okNext = m[i+1]
		count := int(math.Max(float64(vPrev), float64(vNext)))

		if okPrev {
			count = count + 1
		}
		if okNext {
			count = count + 1
		}

		if count > 0 {
			m[i] = count
		} else {
			m[i] = 1
		}
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
