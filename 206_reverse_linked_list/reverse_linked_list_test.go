package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected []int
}

var data = []TestCase{
	{
		nums:     []int{1, 2, 3, 4, 5},
		expected: []int{5, 4, 3, 2, 1},
	},
	{
		nums:     []int{1, 2},
		expected: []int{2, 1},
	},
	{
		nums:     []int{1, 1},
		expected: []int{1, 1},
	},
	{
		nums:     []int{1},
		expected: []int{1},
	},
	{
		nums:     []int{},
		expected: []int{},
	},
}

var solvers = []Solver{
	MutableSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			nums, _ := New(d.nums)
			assert.ElementsMatch(t, d.expected, Walk(s.ReverseList(nums)))
		}
	}
}

func TestCreateAndWalk(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	node, _ := New(arr)
	assert.ElementsMatch(t, arr, Walk(node))
	arr = []int{}
	node, _ = New(arr)
	assert.ElementsMatch(t, arr, Walk(node))
}
