package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	n        int
	expected []int
}

var data = []TestCase{
	{
		nums:     []int{1, 2, 3, 4, 5},
		n:        2,
		expected: []int{1, 2, 3, 5},
	},
	{
		nums:     []int{1, 2},
		n:        1,
		expected: []int{1},
	},
	{
		nums:     []int{1, 2},
		n:        2,
		expected: []int{2},
	},
	{
		nums:     []int{1},
		n:        1,
		expected: []int{},
	},
}

var solvers = []Solver{
	TwoPassSolver{},
	OnePassSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.ElementsMatch(t, d.expected, ToArray(s.RemoveNthFromEnd(New(d.nums), d.n)))
		}
	}
}
