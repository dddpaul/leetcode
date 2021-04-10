package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndWalk(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	node := New(arr)
	walk := node.Walk()
	assert.ElementsMatch(t, arr, walk)
}

// func TestExample1(t *testing.T) {
// 	l1 := []int{2, 4, 3}
// 	l2 := []int{5, 6, 4}
// 	expected := []int{7, 0, 8}

// 	assert.ElementsMatch(t, expected, AddTwoNumbers(l1, l2))
// }
