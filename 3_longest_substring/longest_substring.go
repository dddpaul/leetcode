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

	exists := struct{}{}
	m := make(map[string]interface{})

	i, j, l := 0, 0, 0
	for {
		i = j
		for {
			ch := s[i]
			if _, ok := m[string(ch)]; ok {
				break
			} else {
				m[string(ch)] = exists
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
			delete(m, k)
		}
	}

	return l
}
