package main

/**
 * https://leetcode.com/problems/add-two-numbers/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(nums []int) *ListNode {
	var root, node, prev *ListNode
	for i, v := range nums {
		node = &ListNode{
			Val: v,
		}
		if i == 0 {
			root = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return root
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

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
