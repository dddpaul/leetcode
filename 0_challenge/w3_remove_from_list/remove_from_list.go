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

func Walk(l *ListNode) <-chan *ListNode {
	ch := make(chan *ListNode)
	go func(ch chan *ListNode) {
		defer close(ch)
		for node := l; node != nil; node = node.Next {
			ch <- node
		}
	}(ch)
	return ch
}

func ToArray(l *ListNode) []int {
	s := make([]int, 0)
	for node := range Walk(l) {
		s = append(s, node.Val)
	}
	return s
}

type DirectSolver struct{}

func (s DirectSolver) RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	walk := func(l *ListNode) <-chan *ListNode {
		ch := make(chan *ListNode)
		go func(ch chan *ListNode) {
			defer close(ch)
			for node := l; node != nil; node = node.Next {
				ch <- node
			}
		}(ch)
		return ch
	}

	length := 0
	for range walk(head) {
		length = length + 1
	}

	if length == 1 && n == 1 {
		return nil
	}

	node := head
	var prev *ListNode
	i := 0

	for {
		if prev != nil && i == length-n {
			prev.Next = node.Next
		}
		if node.Next == nil {
			break
		}
		prev = node
		node = node.Next
		i = i + 1
	}
	return head
}
