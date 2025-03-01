package p206

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func reverseList(head *ListNode) *ListNode {
	prev, current := (*ListNode)(nil), head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
