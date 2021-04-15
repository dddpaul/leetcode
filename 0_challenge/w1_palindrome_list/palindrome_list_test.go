package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected bool
}

var data = []TestCase{
	{
		nums:     []int{1, 2, 3, 2, 1},
		expected: true,
	},
	{
		nums:     []int{1, 2, 2, 1},
		expected: true,
	},
	{
		nums:     []int{1, 2, 2, 2, 1},
		expected: true,
	},
	{
		nums:     []int{1, 2, 3, 3, 2, 1},
		expected: true,
	},
	{
		nums:     []int{1, 1},
		expected: true,
	},
	{
		nums:     []int{1, 2, 3, 2, 0},
		expected: false,
	},
	{
		nums:     []int{1, 2},
		expected: false,
	},
	{
		nums:     []int{1, 0, 0},
		expected: false,
	},
	{
		nums:     []int{1, 1, 2, 1},
		expected: false,
	},
	{
		nums:     []int{1},
		expected: true,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			nums, _ := New(d.nums)
			assert.Equal(t, d.expected, s.IsPalindrome(nums))
		}
	}
}
