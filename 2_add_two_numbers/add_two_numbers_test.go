package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	l1       []int
	l2       []int
	expected []int
}

var data = []TestCase{
	{
		l1:       []int{2, 4, 3},
		l2:       []int{5, 6, 4},
		expected: []int{7, 0, 8},
	},
	{
		l1:       []int{0},
		l2:       []int{0},
		expected: []int{0},
	},
	{
		l1:       []int{9, 9, 9, 9, 9, 9, 9},
		l2:       []int{9, 9, 9, 9},
		expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}

func TestCreateAndWalk(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	node, _ := New(arr)
	assert.ElementsMatch(t, arr, node.Walk())
}

func TestExamples(t *testing.T) {
	for _, d := range data {
		l1, _ := New(d.l1)
		l2, _ := New(d.l2)
		assert.ElementsMatch(t, d.expected, AddTwoNumbers(l1, l2).Walk())
	}
}
