package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums            []int
	target          int
	expected        []int
	bisect_expected int
}

var data = []TestCase{
	{
		nums:     []int{1, 1, 1, 1, 2, 2, 4, 4, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 9, 10, 11, 12},
		target:   12,
		expected: []int{22, 22},
	},
	{
		nums:     []int{5, 7, 7, 8, 8, 10},
		target:   8,
		expected: []int{3, 4},
	},
	{
		nums:     []int{5, 7, 7, 8, 8, 10},
		target:   10,
		expected: []int{5, 5},
	},
	{
		nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		target:   1,
		expected: []int{1, 1},
	},
	{
		nums:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		target:   8,
		expected: []int{8, 8},
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
	{
		nums:     []int{1},
		target:   1,
		expected: []int{0, 0},
	},
	{
		nums:     []int{1, 3},
		target:   1,
		expected: []int{0, 0},
	},
	{
		nums:     []int{2, 2},
		target:   3,
		expected: []int{-1, -1},
	},
	{
		nums:     []int{1, 5},
		target:   4,
		expected: []int{-1, -1},
	},
}

var solvers = []Solver{
	DirectSolver{},
	BisectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.ElementsMatch(t, d.expected, s.SearchRange(d.nums, d.target))
		}
	}
}
