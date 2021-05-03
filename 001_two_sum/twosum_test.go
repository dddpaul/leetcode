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
		nums:     []int{2, 7, 11, 15},
		target:   9,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 2, 4},
		target:   6,
		expected: []int{1, 2},
	},
	{
		nums:     []int{3, 3},
		target:   6,
		expected: []int{0, 1},
	},
	{
		nums:     []int{3, 3},
		target:   7,
		expected: []int{},
	},
}

var solvers = []Solver{
	BruteforceSolver{},
	HashSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.ElementsMatch(t, d.expected, s.TwoSum(d.nums, d.target))
		}
	}
}
