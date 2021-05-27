package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	n        int
	pick     int
	expected int
}

var data = []TestCase{
	{
		n:        10,
		pick:     6,
		expected: 6,
	},
	{
		n:        1,
		pick:     1,
		expected: 1,
	},
	{
		n:        2,
		pick:     1,
		expected: 1,
	},
	{
		n:        2,
		pick:     2,
		expected: 2,
	},
}

func TestExamples(t *testing.T) {
	for _, d := range data {
		s := DirectSolver{
			pick: d.pick,
		}
		assert.Equal(t, d.expected, s.GuessNumber(d.n), fmt.Sprintf("For n = %d", d.n))
	}
}
