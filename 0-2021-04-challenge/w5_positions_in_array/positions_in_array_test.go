package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	target   int
	expected []int
}

var data = []TestCase{
	{
		nums:     []int{5, 7, 7, 8, 8, 10},
		target:   8,
		expected: []int{3, 4},
	},
	{
		nums:     []int{5, 7, 7, 8, 8, 10},
		target:   6,
		expected: []int{-1, -1},
	},
	{
		nums:     []int{},
		target:   0,
		expected: []int{-1, -1},
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.ElementsMatch(t, d.expected, s.SearchRange(d.nums, d.target))
		}
	}
}
