package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndWalk(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	node, _ := New(arr)
	assert.ElementsMatch(t, arr, node.Walk())
}

func TestExample1(t *testing.T) {
	l1, _ := New([]int{2, 4, 3})
	l2, _ := New([]int{5, 6, 4})
	expected := []int{7, 0, 8}

	assert.ElementsMatch(t, expected, AddTwoNumbers(l1, l2).Walk())
}

func TestExample2(t *testing.T) {
	l1, _ := New([]int{0})
	l2, _ := New([]int{0})
	expected := []int{0}

	assert.ElementsMatch(t, expected, AddTwoNumbers(l1, l2).Walk())
}

func TestExample3(t *testing.T) {
	l1, _ := New([]int{9, 9, 9, 9, 9, 9, 9})
	l2, _ := New([]int{9, 9, 9, 9})
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	assert.ElementsMatch(t, expected, AddTwoNumbers(l1, l2).Walk())
}
