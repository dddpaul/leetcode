package main

import "errors"

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

func (s DirectSolver) IsPalindrome(head *ListNode) bool {
	return false
}
