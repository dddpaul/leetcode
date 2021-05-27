package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	target   int
	expected int
}

var data = []TestCase{
	{
		nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		target:   1,
		expected: 1,
	},
	{
		nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		target:   8,
		expected: 8,
	},
	{
		nums:     []int{},
		target:   0,
		expected: -1,
	},
	{
		nums:     []int{1},
		target:   1,
		expected: 0,
	},
	{
		nums:     []int{1, 3},
		target:   1,
		expected: 0,
	},
	{
		nums:     []int{2, 2},
		target:   3,
		expected: -1,
	},
	{
		nums:     []int{1, 5},
		target:   4,
		expected: -1,
	},
}

var solvers = []Solver{
	RecursiveSolver{},
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.Search(d.nums, d.target))
		}
	}
}
