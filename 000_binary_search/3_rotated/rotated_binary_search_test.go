package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums      []int
	target    int
	expected  int
	max_index int
}

var data = []TestCase{
	{
		nums:      []int{4, 5, 6, 7, 0, 1, 2},
		target:    0,
		expected:  4,
		max_index: 3,
	},
	{
		nums:      []int{4, 5, 6, 7, 0, 1, 3},
		target:    3,
		expected:  6,
		max_index: 3,
	},
	{
		nums:      []int{4, 5, 6, 7, 0, 1, 2},
		target:    3,
		expected:  -1,
		max_index: 3,
	},
	{
		nums:      []int{6, 7, 0, 1, 2, 3, 4, 5},
		target:    6,
		expected:  0,
		max_index: 1,
	},
	{
		nums:      []int{1, 2, 3, 4, 5, 6, 7, 0},
		target:    6,
		expected:  5,
		max_index: 6,
	},
	{
		nums:      []int{1},
		target:    0,
		expected:  -1,
		max_index: -1,
	},
	{
		nums:      []int{},
		target:    0,
		expected:  -1,
		max_index: -1,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.Search(d.nums, d.target), fmt.Sprintf("For nums = %v, target = %d", d.nums, d.target))
		}
	}
}

func TestFindMaxIndex(t *testing.T) {
	for _, d := range data {
		assert.Equal(t, d.max_index, find_max_index(d.nums), fmt.Sprintf("For nums = %v", d.nums))
	}
}
