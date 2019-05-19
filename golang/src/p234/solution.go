package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	l, r := head, head
	for l.Next != nil && r.Next != nil && r.Next.Next != nil {
		l = l.Next
		r = r.Next.Next
	}

	pre, l := l, l.Next

	for l != nil {
		post := l.Next
		l.Next = pre
		pre = l
		l = post
	}

	l, r = head, pre
	for l != r && l != nil && r != nil {
		if l.Val != r.Val {
			return false
		}
		if l.Next == r && r.Next == l {
			break
		}
		l = l.Next
		r = r.Next
	}
	return true
}

func main() {
	a := &ListNode{1, &ListNode{2, &ListNode{2, nil}}}
	r := isPalindrome(a)
	fmt.Println(r)
}
