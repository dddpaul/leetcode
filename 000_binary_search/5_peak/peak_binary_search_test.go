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
		nums:     []int{1, 2, 3, 1},
		expected: 2,
	},
	{
		nums:     []int{1, 2, 1, 3, 5, 6, 4},
		expected: 5,
	},
	{
		nums:     []int{1},
		expected: 0,
	},
	{
		nums:     []int{1, 0},
		expected: 0,
	},
	{
		nums:     []int{0, 1},
		expected: 1,
	},
}

func TestExamples(t *testing.T) {
	for _, d := range data {
		s := DirectSolver{}
		assert.Equal(t, d.expected, s.FindPeakElement(d.nums), fmt.Sprintf("For nums = %x", d.nums))
	}
}
