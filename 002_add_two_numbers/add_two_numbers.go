package main

import (
	"errors"
)

/**
 * https://leetcode.com/problems/add-two-numbers/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solver interface {
	AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode
}

func New(nums []int) (*ListNode, error) {
	var head, node, prev *ListNode
	for _, v := range nums {
		if v < 0 || v > 9 {
			return nil, errors.New("Non digit")
		}
		node = &ListNode{
			Val: v,
		}
		if prev == nil {
			head = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return head, nil
}

func (l *ListNode) WalkChannel() <-chan *ListNode {
	ch := make(chan *ListNode)
	go func(ch chan *ListNode) {
		defer close(ch)
		node := l
		for {
			ch <- node
			if node.Next == nil {
				break
			}
			node = node.Next
		}
	}(ch)
	return ch
}

func (l *ListNode) Walk() []int {
	arr := make([]int, 0)
	for node := range l.WalkChannel() {
		arr = append(arr, node.Val)
	}
	return arr
}

type DirectSolver struct{}

func (s DirectSolver) AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, node, prev *ListNode
	var sum, overflow, prev_overflow, v1, v2 int
	var node1_walked, node2_walked bool
	node1 := l1
	node2 := l2

	Sum := func(v1 int, v2 int, overflow int) (int, int) {
		sum := v1 + v2 + overflow
		if sum >= 10 {
			return sum - 10, 1
		}
		return sum, 0
	}

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
			head = node
		} else {
			prev.Next = node
		}

		if node1.Next == nil {
			node1_walked = true
			v1 = 0
		} else {
			node1 = node1.Next
		}
		if node2.Next == nil {
			node2_walked = true
			v2 = 0
		} else {
			node2 = node2.Next
		}

		if node1_walked && node2_walked && overflow == 0 {
			break
		}
		prev = node
		prev_overflow = overflow
	}
	return head
}

type ChannelSolver struct{}

func (s ChannelSolver) AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
