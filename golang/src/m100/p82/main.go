package p82

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	var dummy = &ListNode{Next: head, Val: -1 << 31}
	var prev = dummy

	var count = 0
	for node := head; node != nil; node = node.Next {
		if prev.Next.Val == node.Val {
			count++
			prev.Next = node
		} else if count == 1 {
			count = 1
			prev = prev.Next
		} else {
			count = 1
			prev.Next = node
		}
	}
	if count > 1 {
		prev.Next = nil
	}
	return dummy.Next
}
