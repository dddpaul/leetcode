package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/125/template-i/951/
 */

type Solver interface {
	GuessNumber(n int) int
}

type DirectSolver struct {
	pick int
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if guess number is lower than num
 *			      1 if guess number is higher than num
 *               otherwise return 0
 */
func (ds DirectSolver) guess(num int) int {
	if num < ds.pick {
		return 1
	} else if num > ds.pick {
		return -1
	} else {
		return 0
	}
}

func (ds DirectSolver) GuessNumber(n int) int {
	guess := ds.guess
	l, r := 0, n
	for {
		i := l + (r-l)/2
		if guess(i) == 1 {
			l = i + 1
		} else if guess(i) == -1 {
			r = i - 1
		} else {
			return i
		}
		if l > r {
			return i
		}
	}
}
