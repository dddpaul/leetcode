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

func Sum(v1 int, v2 int, overflow int) (int, int) {
	fmt.Printf("%v, %v, %v\n", v1, v2, overflow)
	sum := v1 + v2 + overflow
	if sum >= 10 {
		return sum - 10, 1
	}
	return sum, 0
}

func (l1 *ListNode) SumAll(l2 *ListNode) *ListNode {
	var root, node, prev *ListNode
	var sum, overflow, prev_overflow, v1, v2 int
	var node1_walked, node2_walked bool
	node1 := l1
	node2 := l2
	for {
		if !node1_walked {
			v1 = node1.Val
		}
		if !node2_walked {
			v2 = node2.Val
		}
		sum, overflow = Sum(v1, v2, prev_overflow)
		node = &ListNode{
			Val: sum,
		}
		if prev == nil {
			root = node
		} else {
			prev.Next = node
		}
		if node1.Next == nil {
			node1_walked = true
			v1 = 0
		}
		if node2.Next == nil {
			node2_walked = true
			v2 = 0
		}
		if node1_walked && node2_walked && overflow == 0 {
			break
		}
		prev = node
		prev_overflow = overflow
		if !node1_walked {
			node1 = node1.Next
		}
		if !node2_walked {
			node2 = node2.Next
		}
	}
	return root
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return l1.SumAll(l2)
}
