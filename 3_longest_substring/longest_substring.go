package main

import "fmt"

/**
 * https://leetcode.com/problems/longest-substring-without-repeating-characters/
 */

type Solver interface {
	LengthOfLongestSubstring(s string) int
}

type DirectSolver struct{}

func (ds DirectSolver) LengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	m := make(map[string]int)
	for _, r := range s {
		m[string(r)] = 1
	}

	i, j, l := 0, 0, 0
	for {
		i = j
		for {
			ch := s[i]
			if m[string(ch)] == 1 {
				m[string(ch)] = 0
			} else {
				break
			}
			i = i + 1
			if i == len(s) {
				break
			}
		}
		if l < i-j {
			l = i - j
			fmt.Printf("j = %v, length = %v\n", j, l)
		}
		j = j + 1
		if j == len(s) {
			break
		}
		for k := range m {
			m[k] = 1
		}
	}

	return l
}
