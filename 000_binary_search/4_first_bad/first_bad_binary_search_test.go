package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	n        int
	bad      int
	expected int
}

var data = []TestCase{
	{
		n:        5,
		bad:      4,
		expected: 4,
	},
	{
		n:        1,
		bad:      1,
		expected: 1,
	},
	{
		n:        2,
		bad:      1,
		expected: 1,
	},
	{
		n:        2,
		bad:      2,
		expected: 2,
	},
}

func TestExamples(t *testing.T) {
	for _, d := range data {
		s := DirectSolver{
			bad: d.bad,
		}
		assert.Equal(t, d.expected, s.FirstBadVersion(d.n), fmt.Sprintf("For n = %d", d.n))
	}
}
