package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3712/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solver interface {
	RemoveNthFromEnd(head *ListNode, n int) *ListNode
}

func New(nums []int) *ListNode {
	var head, node, prev *ListNode
	for _, v := range nums {
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
	return head
}

type DirectSolver struct{}

func (s DirectSolver) RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	return nil
}
