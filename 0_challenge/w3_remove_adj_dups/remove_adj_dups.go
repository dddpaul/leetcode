package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3710/
 */

type Solver interface {
	RemoveDuplicates(s string, k int) string
}

type RewindSolver struct{}

func (ds RewindSolver) RemoveDuplicates(s string, k int) string {
	dups := 1
	i := 0
	var prev byte
	if len(s) == 0 || k == 1 {
		return ""
	}
	for {
		ch := s[i]
		if ch == prev {
			dups = dups + 1
		} else {
			dups = 1
		}
		if dups == k {
			// fmt.Printf("i = %d, dups = %d, s = %s", i, dups, s)
			s = s[:i-k+1] + s[i+1:]
			dups = 1
			i = 0
			prev = 0
			// fmt.Printf(", remove = %s, s = %s\n", string(ch), s)
		} else {
			prev = ch
			i = i + 1
			if i == len(s) {
				break
			}
		}
		if len(s) == 0 {
			break
		}
	}
	return s
}

type LinearSolver struct{}

func (ls LinearSolver) RemoveDuplicates(s string, k int) string {
	dups := 1
	i := 0
	var prev byte
	if len(s) == 0 || k == 1 {
		return ""
	}
	for {
		ch := s[i]
		if ch == prev {
			dups = dups + 1
		} else {
			dups = 1
		}
		if dups == k {
			// fmt.Printf("i = %d, dups = %d, s = %s", i, dups, s)
			s = s[:i-k+1] + s[i+1:]
			dups = 1
			i = i - k
			if i > k {
				i = i - k
			} else {
				i = 0
			}
			prev = 0
			// fmt.Printf(", i = %d, remove = %s, s = %s\n", i, string(ch), s)
		} else {
			prev = ch
			i = i + 1
			if i == len(s) {
				break
			}
		}
		if len(s) == 0 {
			break
		}
	}
	return s
}
