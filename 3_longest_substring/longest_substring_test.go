package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	s        string
	expected int
}

var data = []TestCase{
	{
		s:        "abcabcbb",
		expected: 3,
	},
	{
		s:        "bbbbb",
		expected: 1,
	},
	{
		s:        "pwwkew",
		expected: 3,
	},
	{
		s:        "",
		expected: 0,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.LengthOfLongestSubstring(d.s))
		}
	}
}
