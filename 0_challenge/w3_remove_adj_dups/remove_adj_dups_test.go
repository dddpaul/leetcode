package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	s        string
	k        int
	expected string
}

var data = []TestCase{
	{
		s:        "abcd",
		k:        2,
		expected: "abcd",
	},
	{
		s:        "deeedbbcccbdaa",
		k:        3,
		expected: "aa",
	},
	{
		s:        "pbbcggttciiippooaais",
		k:        2,
		expected: "ps",
	},
}

var solvers = []Solver{
	DirectSolver{},
}

func TestExamples(t *testing.T) {
	for _, s := range solvers {
		for _, d := range data {
			assert.Equal(t, d.expected, s.RemoveDuplicates(d.s, d.k))
		}
	}
}
