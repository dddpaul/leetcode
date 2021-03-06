package main

/**
 * https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/595/week-3-april-15th-april-21st/3710/
 */

type Solver interface {
	RemoveDuplicates(s string, k int) string
}

type ArrayRewindSolver struct{}

// n^2 complexity because of array memmove
func (ds ArrayRewindSolver) RemoveDuplicates(s string, k int) string {
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

type ListRewindSolver struct{}

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

func (ls ListRewindSolver) RemoveDuplicates(s string, k int) string {

	if len(s) == 0 || k == 1 {
		return ""
	}

	type ListNode struct {
		Val  byte
		Next *ListNode
		Prev *ListNode
	}

	new := func(s string) *ListNode {
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

	seek := func(l *ListNode, offset int) *ListNode {
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

	walk := func(l *ListNode, forward bool) <-chan *ListNode {
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

	toString := func(l *ListNode, forward bool) string {
		s := make([]byte, 0)
		for node := range walk(l, forward) {
			s = append(s, node.Val)
		}
		return string(s)
	}

	dups := 1
	l := new(s)
	node := l
	var prev *ListNode

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
			// fmt.Printf("dups = %d, s = %s", dups, String(l, true))
			prev = seek(node, -k)
			next := node.Next
			if prev != nil {
				prev.Next = next
			} else {
				l = next
				node = l
			}
			if next != nil {
				next.Prev = prev
			}
			dups = 1
			// fmt.Printf(", remove = %s, s = %s", string(node.Val), String(l, true))
			node = seek(prev, -k)
			if node == nil {
				node = l
				prev = nil
			}
			// fmt.Printf(", seek = %v, prev = %v\n", node, prev)
		}
		if node == nil || node.Next == nil {
			break
		}
		if prev != nil {
			node = node.Next
		}
	}
	return toString(l, true)
}
