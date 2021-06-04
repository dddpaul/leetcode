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
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   0,
		expected: 4,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 3},
		target:   3,
		expected: 6,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   3,
		expected: -1,
	},
	{
		nums:     []int{1},
		target:   0,
		expected: -1,
	},
	{
		nums:     []int{},
		target:   0,
		expected: -1,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.Search(d.nums, d.target))
		}
	}
}
