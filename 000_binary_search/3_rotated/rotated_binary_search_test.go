package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	target   int
	expected int
	pivot    int
}

var data = []TestCase{
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   0,
		expected: 4,
		pivot:    3,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 3},
		target:   3,
		expected: 6,
		pivot:    3,
	},
	{
		nums:     []int{4, 5, 6, 7, 0, 1, 2},
		target:   3,
		expected: -1,
		pivot:    3,
	},
	{
		nums:     []int{6, 7, 0, 1, 2, 3, 4, 5},
		target:   6,
		expected: 0,
		pivot:    1,
	},
	{
		nums:     []int{1, 2, 3, 4, 5, 6, 7, 0},
		target:   6,
		expected: 5,
		pivot:    6,
	},
	{
		nums:     []int{1},
		target:   0,
		expected: -1,
		pivot:    -1,
	},
	{
		nums:     []int{},
		target:   0,
		expected: -1,
		pivot:    -1,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

// func TestExamples(t *testing.T) {
// 	for _, s := range solvers {
// 		for _, d := range data {
// 			assert.Equal(t, d.expected, s.Search(d.nums, d.target))
// 		}
// 	}
// }

func TestFindPivot(t *testing.T) {
	for _, d := range data {
		assert.Equal(t, d.pivot, find_pivot(d.nums), fmt.Sprintf("For nums = %v", d.nums))
	}
}
