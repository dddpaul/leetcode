package main

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
		if _, ok := m[i]; ok {
			continue
		}

		vPrev, okPrev := m[i-1]
		vNext, okNext := m[i+1]
		count := vPrev + vNext + 1

		m[i] = count
		if okPrev {
			m[i-1] = count     // update prev
			m[i-vPrev] = count // update leftmost
		}
		if okNext {
			m[i+1] = count     // update next
			m[i+vNext] = count // update rightmost
		}
	}

	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	return max
}
