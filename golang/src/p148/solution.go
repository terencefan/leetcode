package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func merge(p, q *ListNode) *ListNode {
	head := &ListNode{0, nil}
	node := head
	for {
		if p == nil {
			node.Next = q
			break
		}
		if q == nil {
			node.Next = p
			break
		}
		if q.Val > p.Val {
			node.Next = p
			p = p.Next
		} else {
			node.Next = q
			q = q.Next
		}
		node = node.Next
	}
	return head.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	l, r := head, head
	for l.Next != nil && r.Next != nil && r.Next.Next != nil {
		l = l.Next
		r = r.Next.Next
	}
	r = l.Next
	l.Next = nil

	p := sortList(head)
	q := sortList(r)
	return merge(p, q)
}

func main() {
	root := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	root = sortList(root)
	fmt.Println(root)
}
