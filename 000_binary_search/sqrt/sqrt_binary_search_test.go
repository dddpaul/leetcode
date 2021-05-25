package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	x        int
	expected int
}

var data = []TestCase{
	{
		x:        8,
		expected: 2,
	},
	{
		x:        7,
		expected: 2,
	},
	{
		x:        6,
		expected: 2,
	},
	{
		x:        5,
		expected: 2,
	},
	{
		x:        4,
		expected: 2,
	},
	{
		x:        3,
		expected: 1,
	},
	{
		x:        1,
		expected: 1,
	},
	{
		x:        0,
		expected: 0,
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.MySqrt(d.x), fmt.Sprintf("For x = %d", d.x))
		}
	}
}
