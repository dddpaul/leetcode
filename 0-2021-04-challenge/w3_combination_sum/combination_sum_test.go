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
		nums:     []int{1, 2, 3},
		target:   4,
		expected: 7,
	},
	{
		nums:     []int{9},
		target:   3,
		expected: 0,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.CombinationSum4(d.nums, d.target))
		}
	}
}
