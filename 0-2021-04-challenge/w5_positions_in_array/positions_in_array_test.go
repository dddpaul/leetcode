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
		nums:            []int{5, 7, 7, 8, 8, 10},
		target:          8,
		expected:        []int{3, 4},
		bisect_expected: 3,
	},
	{
		nums:            []int{5, 7, 7, 8, 8, 10},
		target:          10,
		expected:        []int{5, 5},
		bisect_expected: 5,
	},
	{
		nums:            []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		target:          1,
		expected:        []int{1, 1},
		bisect_expected: 1,
	},
	{
		nums:            []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		target:          8,
		expected:        []int{8, 8},
		bisect_expected: 8,
	},
	{
		nums:            []int{5, 7, 7, 8, 8, 10},
		target:          6,
		expected:        []int{-1, -1},
		bisect_expected: -1,
	},
	{
		nums:            []int{},
		target:          0,
		expected:        []int{-1, -1},
		bisect_expected: -1,
	},
	{
		nums:            []int{1},
		target:          1,
		expected:        []int{0, 0},
		bisect_expected: 0,
	},
	{
		nums:            []int{1, 3},
		target:          1,
		expected:        []int{0, 0},
		bisect_expected: 0,
	},
	{
		nums:            []int{2, 2},
		target:          3,
		expected:        []int{-1, -1},
		bisect_expected: -1,
	},
	{
		nums:            []int{1, 5},
		target:          4,
		expected:        []int{-1, -1},
		bisect_expected: -1,
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

func TestBisect(t *testing.T) {
	bs := BisectSolver{}
	for _, d := range data {
		assert.Equal(t, d.bisect_expected, bs.SearchRange1(d.nums, d.target))
	}
}
