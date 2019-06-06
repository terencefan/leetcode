package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre, post *ListNode
	for head != nil {
		post = head.Next
		head.Next = pre
		pre = head
		head = post
	}
	return pre
}
