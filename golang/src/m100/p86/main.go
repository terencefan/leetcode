package p86

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func partition(head *ListNode, x int) *ListNode {

	var d1 = &ListNode{}
	var d2 = &ListNode{}
	l1, l2 := d1, d2

	for node := head; node != nil; {
		next := node.Next
		if node.Val < x {
			l1.Next, l1 = node, node
			l1.Next = nil
		} else {
			l2.Next, l2 = node, node
			l2.Next = nil
		}
		node = next
	}

	if d1.Next != nil {
		l1.Next = d2.Next
		return d1.Next
	} else {
		return d2.Next
	}
}
