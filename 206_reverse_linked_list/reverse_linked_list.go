package main

/*
 * https://leetcode.com/problems/reverse-linked-list/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solver interface {
	ReverseList(head *ListNode) *ListNode
}

func New(nums []int) (*ListNode, error) {
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
	return head, nil
}

func WalkChannel(l *ListNode) <-chan *ListNode {
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

func Walk(l *ListNode) []int {
	arr := make([]int, 0)
	for node := range WalkChannel(l) {
		arr = append(arr, node.Val)
	}
	return arr
}

type MutableSolver struct{}

func (s MutableSolver) ReverseList(head *ListNode) *ListNode {
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
