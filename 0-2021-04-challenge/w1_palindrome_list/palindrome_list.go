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
			if node == nil {
				return
			}
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

	reverse := func(head *ListNode) *ListNode {
		var prev *ListNode
		node := head
		if node == nil {
			return nil
		}
		for {
			next := node.Next
			if prev == nil {
				node.Next = nil
			} else {
				node.Next = prev
			}
			prev = node
			if next == nil {
				break
			}
			node = next
		}
		return prev
	}

	length := 0
	for range walk(head) {
		length = length + 1
	}
	if length == 1 {
		return true
	}

	half := head
	for i := 0; i < length/2+length%2; i++ {
		half = half.Next
	}

	n1 := head
	n2 := reverse(half)

	for {
		if n1.Val != n2.Val {
			return false
		}
		if n2.Next == nil {
			return true
		}
		n1 = n1.Next
		n2 = n2.Next
	}
}
