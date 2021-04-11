package main

import (
	"errors"
	"fmt"
)

/**
 * https://leetcode.com/problems/add-two-numbers/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(nums []int) (*ListNode, error) {
	var root, node, prev *ListNode
	for _, v := range nums {
		if v < 0 || v > 9 {
			return nil, errors.New("Non digit")
		}
		node = &ListNode{
			Val: v,
		}
		if prev == nil {
			root = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return root, nil
}

func (l *ListNode) Walk() []int {
	arr := make([]int, 0)
	node := l
	for {
		arr = append(arr, node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return arr
}

func (l1 *ListNode) Sum(l2 *ListNode) (int, int) {
	fmt.Printf("%v, %v\n", l1.Val, l2.Val)
	sum := l1.Val + l2.Val
	if sum >= 10 {
		return sum - 10, 1
	}
	return sum, 0
}

func (l1 *ListNode) SumAll(l2 *ListNode) *ListNode {
	var root, node, prev *ListNode
	var sum, overflow, prev_overflow int
	node1 := l1
	node2 := l2
	for {
		sum, overflow = node1.Sum(node2)
		node = &ListNode{
			Val: sum + prev_overflow,
		}
		if prev == nil {
			root = node
		} else {
			prev.Next = node
		}
		if node1.Next == nil {
			break
		}
		prev = node
		prev_overflow = overflow
		node1 = node1.Next
		node2 = node2.Next
	}
	return root
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return l1.SumAll(l2)
}
