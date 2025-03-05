package p25

import (
	"terencefan.com/leetcode/src/utils"
)

type ListNode = utils.ListNode

func reverse(left, right *ListNode) {
	var prev *ListNode = nil

	for left != right {
		next := left.Next
		if prev != nil {
			left.Next = prev
		}
		prev, left = left, next
	}
	left.Next = prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}

	var dummy = &ListNode{Next: head}
	var node = head
	var groupHead = node
	var groupTail = dummy

	for n := 0; node != nil; n++ {
		next := node.Next
		if n%k == k-1 {
			reverse(groupHead, node)
			groupHead.Next, groupTail.Next = next, node
			groupTail = groupHead
		} else if n%k == 0 {
			groupHead = node
		}
		node = next
	}
	return dummy.Next
}
