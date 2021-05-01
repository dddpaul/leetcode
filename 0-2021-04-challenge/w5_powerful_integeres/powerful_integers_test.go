package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	x, y, bound int
	expected    []int
}

var data = []TestCase{
	{
		x:        2,
		y:        3,
		bound:    10,
		expected: []int{2, 3, 4, 5, 7, 9, 10},
	},
	{
		x:        3,
		y:        5,
		bound:    15,
		expected: []int{2, 4, 6, 8, 10, 14},
	},
	{
		x:        2,
		y:        1,
		bound:    10,
		expected: []int{9, 2, 3, 5},
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.ElementsMatch(t, d.expected, s.PowerfulIntegers(d.x, d.y, d.bound))
		}
	}
}
