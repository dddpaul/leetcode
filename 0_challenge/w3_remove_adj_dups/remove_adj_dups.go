package main

import "fmt"

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3710/
 */

type Solver interface {
	RemoveDuplicates(s string, k int) string
}

type RewindSolver struct{}

// n^2 complexity because of array memmove
func (ds RewindSolver) RemoveDuplicates(s string, k int) string {
	dups := 1
	i := 0
	var prev byte
	if len(s) == 0 || k == 1 {
		return ""
	}
	for {
		ch := s[i]
		if ch == prev {
			dups = dups + 1
		} else {
			dups = 1
		}
		if dups == k {
			// fmt.Printf("i = %d, dups = %d, s = %s", i, dups, s)
			s = s[:i-k+1] + s[i+1:]
			dups = 1
			i = i - k
			if i > k {
				i = i - k
			} else {
				i = 0
			}
			prev = 0
			// fmt.Printf(", i = %d, remove = %s, s = %s\n", i, string(ch), s)
		} else {
			prev = ch
			i = i + 1
			if i == len(s) {
				break
			}
		}
		if len(s) == 0 {
			break
		}
	}
	return s
}

type LinearSolver struct{}

type ListNode struct {
	Val  byte
	Next *ListNode
	Prev *ListNode
}

func New(s string) *ListNode {
	var head, node, prev *ListNode
	for _, v := range s {
		node = &ListNode{
			Val: byte(v),
		}
		if prev == nil {
			head = node
		} else {
			prev.Next = node
			node.Prev = prev
		}
		prev = node
	}
	return head
}

func Walk(l *ListNode, forward bool) <-chan *ListNode {
	ch := make(chan *ListNode)
	go func(ch chan *ListNode) {
		defer close(ch)
		if forward {
			for node := l; node != nil; node = node.Next {
				ch <- node
			}
		} else {
			for node := l; node != nil; node = node.Prev {
				ch <- node
			}
		}
	}(ch)
	return ch
}

func String(l *ListNode, forward bool) string {
	s := make([]byte, 0)
	for node := range Walk(l, forward) {
		s = append(s, node.Val)
	}
	return string(s)
}

func Seek(l *ListNode, offset int) *ListNode {
	i := 0
	node := l
	for {
		if node == nil {
			return nil
		}
		if i < offset {
			node = node.Next
			i = i + 1
		} else if i > offset {
			node = node.Prev
			i = i - 1
		} else {
			return node
		}
	}
}

func (ls LinearSolver) RemoveDuplicates(s string, k int) string {
	if len(s) == 0 || k == 1 {
		return ""
	}
	dups := 1
	l := New(s)
	node := l
	var prev *ListNode
	n := 0
	for {
		if prev != nil {
			if node.Val == prev.Val {
				dups = dups + 1
			} else {
				dups = 1
			}
		}
		prev = node
		if dups == k {
			fmt.Printf("n = %d, dups = %d, s = %s", n, dups, String(l, true))
			prev = Seek(node, -k)
			next := node.Next
			if prev != nil {
				prev.Next = next
			} else {
				l = next
				node = l
			}
			next.Prev = prev
			dups = 1
			fmt.Printf(", remove = %s, s = %s", string(node.Val), String(l, true))
			node = Seek(prev, -k+2)
			if node == nil {
				node = l
				prev = nil
			}
			n = n + 1
			if n == 10 {
				break
			}
			fmt.Printf(", seek = %v, prev = %v\n", node, prev)
		}
		if node.Next == nil {
			break
		}
		if prev != nil {
			node = node.Next
		}
	}
	return String(l, true)
}
