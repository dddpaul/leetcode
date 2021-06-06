package main

/**
 * https://leetcode.com/explore/learn/card/binary-search/126/template-ii/947/
 */

type Solver interface {
	FirstBadVersion(n int) int
}

type DirectSolver struct {
	bad int
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func (ds DirectSolver) isBadVersion(version int) bool {
	if version >= ds.bad {
		return true
	}
	return false
}

func (ds DirectSolver) FirstBadVersion(n int) int {
	isBadVersion := ds.isBadVersion
	if n <= 1 {
		return n
	}
	l, r := 1, n
	for {
		i := l + (r-l)/2
		if l == r {
			return i
		}
		if isBadVersion(i) {
			r = i
		} else {
			l = i + 1
		}
	}
}
