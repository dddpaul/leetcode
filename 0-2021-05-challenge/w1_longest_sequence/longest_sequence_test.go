package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	nums     []int
	expected int
}

var data = []TestCase{
	{
		nums:     []int{100, 4, 200, 1, 3, 2},
		expected: 4,
	},
	{
		nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		expected: 9,
	},
	{
		nums:     []int{1, 2, 0, 1},
		expected: 3,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.LongestConsecutive(d.nums), fmt.Sprintf("For nums = %x", d.nums))
		}
	}
}
