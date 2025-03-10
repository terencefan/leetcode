package p61

import "terencefan.com/leetcode/src/utils"

type ListNode = utils.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	node := head

	n := 0
	for node != nil {
		node = node.Next
		n++
	}

	k = k % n
	if k == 0 {
		return head
	}

	var prev *ListNode
	node = head
	for range n - k {
		prev = node
		node = node.Next
	}
	var r = node
	prev.Next = nil
	prev = nil

	for node != nil {
		prev = node
		node = node.Next
	}
	prev.Next = head
	return r
}
