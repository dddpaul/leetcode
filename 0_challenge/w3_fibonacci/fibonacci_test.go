package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	n        int
	expected int
}

var data = []TestCase{
	{
		n:        0,
		expected: 0,
	},
	{
		n:        1,
		expected: 1,
	},
	{
		n:        2,
		expected: 1,
	},
	{
		n:        3,
		expected: 2,
	},
	{
		n:        4,
		expected: 3,
	},
	{
		n:        5,
		expected: 5,
	},
	{
		n:        10,
		expected: 55,
	},
	{
		n:        20,
		expected: 6765,
	},
}

var solvers = []Solver{
	LinearSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.Fib(d.n))
		}
	}
}
