package main

import (
	"errors"
)

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/593/week-1-april-1st-april-7th/3693/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solver interface {
	IsPalindrome(head *ListNode) bool
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

type DirectSolver struct{}

func (s DirectSolver) IsPalindrome(head *ListNode) bool {

	walk := func(l *ListNode) <-chan *ListNode {
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

	stack := make([]int, 0)
	backward := false
	count := 0

	for node := range walk(head) {
		if count == 0 {
			stack = append(stack, node.Val)
		} else {
			prev := stack[len(stack)-1]
			if node.Val != prev {
				stack = append(stack, node.Val)
				if len(stack) >= 3 {
					prev_prev := stack[len(stack)-3]
					if node.Val == prev_prev {
						stack = stack[:len(stack)-3]
						backward = true
					}
				}
			} else if backward {
				stack = stack[:len(stack)-1]
			}
		}
		count = count + 1
	}

	return count == 1 || (count == 2 && len(stack) == 1) || len(stack) == 0
}
