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
		nums:     []int{1, 2, 2, 1},
		expected: true,
	},
	{
		nums:     []int{1, 2},
		expected: false,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestCreateAndWalk(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	node, _ := New(arr)
	assert.ElementsMatch(t, arr, node.Walk())
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			nums, _ := New(d.nums)
			assert.Equal(t, d.expected, s.IsPalindrome(nums))
		}
	}
}
